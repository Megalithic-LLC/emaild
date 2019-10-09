package cloudservice

import (
	"errors"
	"time"

	"github.com/docktermj/go-logger/logger"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

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
