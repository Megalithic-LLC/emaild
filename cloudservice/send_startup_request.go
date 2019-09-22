package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/agentstreamproto"
)

func (self *CloudService) SendStartupRequest() (*agentstreamproto.ServerMessage, error) {
	req := agentstreamproto.ClientMessage{
		MessageType: &agentstreamproto.ClientMessage_StartupRequest{
			StartupRequest: &agentstreamproto.StartupRequest{},
		},
	}
	return self.SendRequest(req)
}
