package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/drauschenbach/megalithicd/propertykey"
)

func (self *CloudService) handleClaimRequest(requestId uint64, claimReq agentstreamproto.ClaimRequest) {
	logger.Tracef("CloudService:handleClaimRequest(%d)", requestId)

	if err := self.propertiesDAO.Set(propertykey.Token, claimReq.Token); err != nil {
		logger.Errorf("Failed storing token: %v", err)
		self.SendErrorResponse(requestId, err)
		return
	}

	if err := self.SendAckResponse(requestId); err != nil {
		logger.Errorf("Failed sending ack response: %v", err)
	}
}
