//go:generate protoc --proto_path=emailproto --go_out=plugins=grpc:emailproto email.proto
package cloudservice

import (
	"errors"
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
	"github.com/on-prem-net/emaild/propertykey"
)

const API_URL = "API_URL"

type CloudService struct {
	accountsDAO     dao.AccountsDAO
	agentID         string
	cloudServiceURL url.URL
	conn            *websocket.Conn
	db              *genji.DB
	mutex           sync.Mutex
	nextID          uint64
	pending         map[uint64]*Call
	propertiesDAO   dao.PropertiesDAO
	snapshotsDAO    dao.SnapshotsDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	propertiesDAO dao.PropertiesDAO,
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

	self := CloudService{
		accountsDAO:     accountsDAO,
		agentID:         agentID,
		cloudServiceURL: url.URL{Scheme: scheme, Host: parsedURL.Host, Path: "/v1/agentStream"},
		db:              db,
		nextID:          1,
		pending:         map[uint64]*Call{},
		propertiesDAO:   propertiesDAO,
		snapshotsDAO:    snapshotsDAO,
	}

	go self.dialer()
	go self.reader()

	return &self
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

func (self *CloudService) SendRequest(req emailproto.ClientMessage) (*emailproto.ServerMessage, error) {
	self.mutex.Lock()
	req.Id = self.getNextID()
	call := NewCall(req)
	self.pending[req.Id] = call

	// Encode request
	rawMessage, err := proto.Marshal(&req)
	if err != nil {
		logger.Errorf("Failed encoding request: %v", err)
		defer self.mutex.Unlock()
		return nil, err
	}

	// Send request
	if err := self.conn.WriteMessage(websocket.BinaryMessage, rawMessage); err != nil {
		delete(self.pending, req.Id)
		defer self.mutex.Unlock()
		return nil, err
	}

	self.mutex.Unlock()

	select {
	case <-call.Done:
	case <-time.After(2 * time.Second):
		call.Error = errors.New("request timeout")
	}

	return call.Res, call.Error
}

func (self *CloudService) SendResponse(res emailproto.ClientMessage) error {
	logger.Tracef("CloudService:SendResponse()")

	// Encode response
	rawMessage, err := proto.Marshal(&res)
	if err != nil {
		logger.Errorf("Failed encoding response: %v", err)
		return err
	}

	// Send response
	return self.conn.WriteMessage(websocket.BinaryMessage, rawMessage)
}
