package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetEndpointsRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetEndpointsRequest{
			GetEndpointsRequest: &emailproto.GetEndpointsRequest{},
		},
	}
	return self.SendRequest(req)
}
