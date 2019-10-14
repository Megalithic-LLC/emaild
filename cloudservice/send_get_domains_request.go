package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetDomainsRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetDomainsRequest{
			GetDomainsRequest: &emailproto.GetDomainsRequest{},
		},
	}
	return self.SendRequest(req)
}
