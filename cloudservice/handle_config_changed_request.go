package cloudservice

import (
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/drauschenbach/megalithicd/propertykey"
)

func (self *CloudService) handleConfigChangedRequest(requestId uint64, configChangedReq agentstreamproto.ConfigChangedRequest) {
	logger.Tracef("CloudService:handleConfigChangedRequest(%d)", requestId)

	if err := self.SendAckResponse(requestId); err != nil {
		logger.Errorf("Failed sending ack response: %v", err)
	}

	for table, hash := range configChangedReq.HashesByTable {
		hashAsHex := fmt.Sprintf("%x", hash)

		key := fmt.Sprintf(propertykey.HashByTablePattern, table)
		if value, err := self.propertiesDAO.Get(key); err != nil {
			logger.Errorf("Failed looking up table config hash: %v", err)
		} else {
			if hashAsHex == value {
				continue
			}
			self.propertiesDAO.Set(key, hashAsHex)
		}
	}

}
