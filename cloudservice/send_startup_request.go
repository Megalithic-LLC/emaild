package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendStartupRequest() (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_StartupRequest{
			StartupRequest: &emailproto.StartupRequest{
				ServiceId: "blmkmfd5jj89vu275l5g",
			},
		},
	}
	return self.SendRequest(req)
}
