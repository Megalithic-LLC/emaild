package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetSnapshotsRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetSnapshotsRequest{
			GetSnapshotsRequest: &emailproto.GetSnapshotsRequest{},
		},
	}
	return self.SendRequest(req)
}
