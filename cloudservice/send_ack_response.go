package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/agentstreamproto"
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
