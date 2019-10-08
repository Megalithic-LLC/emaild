package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendStartupRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_StartupRequest{
			StartupRequest: &emailproto.StartupRequest{},
		},
	}
	return self.SendRequest(req)
}
