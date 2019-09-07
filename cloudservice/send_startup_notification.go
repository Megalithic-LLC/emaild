package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/drauschenbach/megalithicd/propertykey"
)

type StartupNotificationRequest struct {
	Command string `json:"command"`
	NodeID  string `json:"nodeid"`
}

func (self *CloudService) SendStartupNotification() error {
	nodeID, err := self.propertiesDAO.Get(propertykey.NodeID)
	if err != nil {
		logger.Fatalf("Failed looking up node id: %v", err)
		return nil
	}
	//	req := agentstreamproto.StartupRequest{
	//		NodeId: nodeID,
	//	}
	req := agentstreamproto.ClientMessage{
		MessageType: &agentstreamproto.ClientMessage_StartupRequest{
			StartupRequest: &agentstreamproto.StartupRequest{
				NodeId: nodeID,
			},
		},
	}
	_, err = self.SendRequest(req)
	return err
}
