package cloudservice

import (
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
)

func (self *CloudService) SendStartupNotification() error {
	req := agentstreamproto.ClientMessage{
		MessageType: &agentstreamproto.ClientMessage_StartupRequest{
			StartupRequest: &agentstreamproto.StartupRequest{},
		},
	}
	_, err := self.SendRequest(req)
	return err
}
