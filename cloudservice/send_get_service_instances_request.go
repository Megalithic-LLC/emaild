package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetServiceInstancesRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetServiceInstancesRequest{
			GetServiceInstancesRequest: &emailproto.GetServiceInstancesRequest{},
		},
	}
	return self.SendRequest(req)
}
