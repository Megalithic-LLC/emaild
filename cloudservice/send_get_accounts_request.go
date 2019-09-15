package cloudservice

import (
	"errors"
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
)

func (self *CloudService) SendGetAccountsRequest() (*agentstreamproto.EmailcdnGetAccountsResponse, error) {
	logger.Tracef("CloudService:SendGetAccountsRequest()")

	req := &agentstreamproto.ClientMessage{
		MessageType: &agentstreamproto.ClientMessage_EmailcdnGetAccountsRequest{
			EmailcdnGetAccountsRequest: &agentstreamproto.EmailcdnGetAccountsRequest{},
		},
	}

	res, err := self.SendRequest(req)
	if err != nil {
		return nil, err
	}

	switch res.MessageType.(type) {
	case *agentstreamproto.ServerMessage_EmailcdnGetAccountsResponse:
		return res.GetEmailcdnGetAccountsResponse(), nil
	default:
		return nil, errors.New(fmt.Sprintf("Unexpected response type: %+v", err))
	}

}
