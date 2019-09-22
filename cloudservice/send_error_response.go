package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/agentstreamproto"
)

func (self *CloudService) SendErrorResponse(requestId uint64, err error) error {
	errorRes := agentstreamproto.ClientMessage{
		Id: requestId,
		MessageType: &agentstreamproto.ClientMessage_ErrorResponse{
			ErrorResponse: &agentstreamproto.ErrorResponse{
				Error: err.Error(),
			},
		},
	}
	return self.SendResponse(errorRes)
}
