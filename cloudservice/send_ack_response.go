package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendAckResponse(requestId uint64) error {
	ackRes := emailproto.ClientMessage{
		Id: requestId,
		MessageType: &emailproto.ClientMessage_AckResponse{
			AckResponse: &emailproto.AckResponse{},
		},
	}
	return self.SendResponse(ackRes)
}
