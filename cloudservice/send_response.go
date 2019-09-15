//go:generate protoc --proto_path=agentstreamproto --go_out=plugins=grpc:agentstreamproto agentstream.proto
package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func (self *CloudService) SendResponse(res agentstreamproto.ClientMessage) error {
	// Encode response
	rawMessage, err := proto.Marshal(&res)
	if err != nil {
		logger.Errorf("Failed encoding response: %v", err)
		return err
	}

	// Send response
	return self.conn.WriteMessage(websocket.BinaryMessage, rawMessage)
}
