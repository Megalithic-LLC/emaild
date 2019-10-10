//go:generate protoc --proto_path=emailproto --go_out=plugins=grpc:emailproto email.proto
package cloudservice

import (
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/model"
	"github.com/on-prem-net/emaild/propertykey"
	"github.com/on-prem-net/emaild/snapshotmanager"
)

const API_URL = "API_URL"

type CloudService struct {
	accountsDAO         dao.AccountsDAO
	agentID             string
	cloudServiceURL     url.URL
	conn                *websocket.Conn
	db                  *genji.DB
	domainsDAO          dao.DomainsDAO
	endpointsDAO        dao.EndpointsDAO
	mutex               sync.Mutex
	nextID              uint64
	pending             map[uint64]*Call
	propertiesDAO       dao.PropertiesDAO
	serviceInstancesDAO dao.ServiceInstancesDAO
	snapshotManager     *snapshotmanager.SnapshotManager
	snapshotsDAO        dao.SnapshotsDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	domainsDAO dao.DomainsDAO,
	endpointsDAO dao.EndpointsDAO,
	propertiesDAO dao.PropertiesDAO,
	serviceInstancesDAO dao.ServiceInstancesDAO,
	snapshotManager *snapshotmanager.SnapshotManager,
	snapshotsDAO dao.SnapshotsDAO,
) *CloudService {
	cloudServiceURL := os.Getenv(API_URL)
	if cloudServiceURL == "" {
		cloudServiceURL = "https://api.on-prem.net"
	}
	parsedURL, err := url.Parse(cloudServiceURL)
	if err != nil {
		logger.Fatalf("Malformed %s", API_URL)
		return nil
	}

	agentID, err := propertiesDAO.Get(propertykey.NodeID)
	if err != nil {
		logger.Fatalf("Failed looking up agent id: %v", err)
		return nil

	}

	scheme := "ws"
	if parsedURL.Scheme == "https" {
		scheme = "wss"
	}

	self := &CloudService{
		accountsDAO:         accountsDAO,
		agentID:             agentID,
		cloudServiceURL:     url.URL{Scheme: scheme, Host: parsedURL.Host, Path: "/v1/agentStream"},
		db:                  db,
		domainsDAO:          domainsDAO,
		endpointsDAO:        endpointsDAO,
		nextID:              1,
		pending:             map[uint64]*Call{},
		propertiesDAO:       propertiesDAO,
		serviceInstancesDAO: serviceInstancesDAO,
		snapshotManager:     snapshotManager,
		snapshotsDAO:        snapshotsDAO,
	}

	go self.dialer()
	go self.reader()

	snapshotManager.RegisterListener(self)

	return self
}

func NewCall(req emailproto.ClientMessage) *Call {
	return &Call{
		Req:  req,
		Done: make(chan bool),
	}
}

func (self *CloudService) dialer() {
	for {
		// Already connected?
		if self.conn != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Configure connect request
		token, err := self.propertiesDAO.Get(propertykey.Token)
		if err != nil {
			logger.Fatalf("Failed looking up token: %v", err)
			return
		}
		logger.Debugf("Connecting to %s", self.cloudServiceURL.String())
		header := http.Header{
			"X-AgentID": []string{self.agentID},
		}
		if token != "" {
			header["Authorization"] = []string{"Bearer " + token}
		}

		// Connect
		conn, _, err := websocket.DefaultDialer.Dial(self.cloudServiceURL.String(), header)
		if err != nil {
			self.conn = nil
			time.Sleep(5 * time.Second)
			continue
		}
		logger.Debugf("Connected to %s", self.cloudServiceURL.String())
		self.conn = conn

		// Initial handshake
		if res, err := self.SendStartupRequest(); err != nil {
			logger.Warnf("Failed contacting cloud service: %v", err)
		} else if startupRes := res.GetStartupResponse(); startupRes != nil {
			self.processConfigChanges(startupRes.ConfigHashesByTable)
		}
	}
}

func (self *CloudService) Disconnect() {
	logger.Tracef("CloudService:Disconnect()")
	self.mutex.Lock()
	defer self.mutex.Unlock()
	if self.conn != nil {
		self.conn.Close()
		self.conn = nil
		self.pending = map[uint64]*Call{}
	}
}

func (self *CloudService) getNextID() uint64 {
	id := self.nextID
	self.nextID++
	return id
}

func (self *CloudService) reader() {
	var err error
	for err == nil {
		if self.conn == nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Read message
		_, rawMessage, err := self.conn.ReadMessage()
		if err != nil {
			logger.Errorf("Failed reading message: %v", err)
			break
		}

		// Decode message
		var message emailproto.ServerMessage
		if err := proto.Unmarshal(rawMessage, &message); err != nil {
			logger.Errorf("Failed decoding message: %v", err)
			break
		}

		self.mutex.Lock()
		call, isResponse := self.pending[message.Id]
		if isResponse {
			delete(self.pending, message.Id)
			self.mutex.Unlock()
			call.Res = &message
			call.Done <- true
		} else {
			self.mutex.Unlock()
			self.route(message)
		}
	}

	// Terminate all calls
	self.mutex.Lock()
	for _, call := range self.pending {
		call.Error = err
		call.Done <- true
	}
	self.mutex.Unlock()
}

func (self *CloudService) UpdateSnapshot(snapshot *model.Snapshot) {
	logger.Tracef("CloudService:UpdateSnapshot()")
	if _, err := self.SendUpdateSnapshotRequest(snapshot); err != nil {
		logger.Errorf("Failed updating cloud service with snapshot progress: %v", err)
		return
	}

	// Upload to cloud
	res, err := self.SendGetSnapshotChunksMissingRequest(snapshot.Id)
	if err != nil {
		logger.Errorf("Failed asking cloud service for missing chunks of snapshot: %v", err)
		return
	}
	if getSnapshotChunksMissingResponse := res.GetGetSnapshotChunksMissingResponse(); getSnapshotChunksMissingResponse != nil {
		for _, chunkNumber := range getSnapshotChunksMissingResponse.Chunks {
			data, err := self.snapshotManager.GetChunk(snapshot, chunkNumber)
			if err != nil {
				logger.Errorf("Failed reading snapshot chunk: %v", err)
				return
			}
			self.SendSetSnapshotChunkRequest(snapshot.Id, chunkNumber, data)
		}
		logger.Infof("Snapshot %s uploaded to cloud", snapshot.Id)
	}
}
