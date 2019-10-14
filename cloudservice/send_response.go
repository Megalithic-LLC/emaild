package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

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
