package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetAccountsRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetAccountsRequest{
			GetAccountsRequest: &emailproto.GetAccountsRequest{},
		},
	}
	return self.SendRequest(req)
}
