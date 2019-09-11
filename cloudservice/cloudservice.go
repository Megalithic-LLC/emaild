//go:generate protoc --proto_path=agentstreamproto --go_out=plugins=grpc:agentstreamproto agentstream.proto
package cloudservice

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/drauschenbach/megalithicd/dao"
	"github.com/drauschenbach/megalithicd/propertykey"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

const CLOUDSERVICE_URL = "CLOUDSERVICE_URL"

type CloudService struct {
	agentID         string
	cloudServiceURL url.URL
	conn            *websocket.Conn
	mutex           sync.Mutex
	nextID          uint64
	pending         map[uint64]*Call
	propertiesDAO   dao.PropertiesDAO
}

func New(propertiesDAO dao.PropertiesDAO) *CloudService {
	cloudServiceURL := os.Getenv(CLOUDSERVICE_URL)
	if cloudServiceURL == "" {
		cloudServiceURL = "http://localhost:3000"
	}
	parsedURL, err := url.Parse(cloudServiceURL)
	if err != nil {
		logger.Fatalf("Malformed %s", CLOUDSERVICE_URL)
		return nil
	}

	agentID, err := propertiesDAO.Get(propertykey.NodeID)
	if err != nil {
		logger.Fatalf("Failed looking up agent id: %v", err)
		return nil

	}
	self := CloudService{
		agentID:         agentID,
		cloudServiceURL: url.URL{Scheme: "ws", Host: parsedURL.Host, Path: "/v1/agentStream"},
		pending:         map[uint64]*Call{},
		propertiesDAO:   propertiesDAO,
		nextID:          1,
	}

	go self.dialer()
	go self.reader()

	return &self
}

func NewCall(req agentstreamproto.ClientMessage) *Call {
	return &Call{
		Req:  req,
		Done: make(chan bool),
	}
}

func (self *CloudService) dialer() {
	for {
		if self.conn != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		logger.Debugf("Connecting to %s", self.cloudServiceURL.String())
		header := http.Header{
			"X-AgentID": []string{self.agentID},
		}
		conn, _, err := websocket.DefaultDialer.Dial(self.cloudServiceURL.String(), header)
		if err != nil {
			self.conn = nil
			time.Sleep(5 * time.Second)
			continue
		}
		logger.Debugf("Connected to %s", self.cloudServiceURL.String())
		self.conn = conn

		if _, err := self.SendStartupRequest(); err != nil {
			logger.Warnf("Failed contacting cloud service: %v", err)
		}
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
		var message agentstreamproto.ServerMessage
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

func (self *CloudService) SendRequest(req agentstreamproto.ClientMessage) (*agentstreamproto.ServerMessage, error) {
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

func (self *CloudService) SendResponse(res agentstreamproto.ClientMessage) error {
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
