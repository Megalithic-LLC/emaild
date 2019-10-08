package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/agentstreamproto"
)

func (self *CloudService) SendAckResponse(requestId uint64) error {
	ackRes := agentstreamproto.ClientMessage{
		Id: requestId,
		MessageType: &agentstreamproto.ClientMessage_AckResponse{
			AckResponse: &agentstreamproto.AckResponse{},
		},
	}
	return self.SendResponse(ackRes)
}
