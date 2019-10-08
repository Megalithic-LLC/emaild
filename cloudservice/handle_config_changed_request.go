package cloudservice

import (
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
	"github.com/on-prem-net/emaild/propertykey"
)

func (self *CloudService) handleConfigChangedRequest(requestId uint64, configChangedReq emailproto.ConfigChangedRequest) {
	logger.Tracef("CloudService:handleConfigChangedRequest(%d)", requestId)

	if err := self.SendAckResponse(requestId); err != nil {
		logger.Errorf("Failed sending ack response: %v", err)
	}

	self.processConfigChanges(configChangedReq.HashesByTable)
}

func (self *CloudService) processConfigChanges(configHashesByTable map[string][]byte) {
	logger.Tracef("CloudService:processConfigChanges()")
	for table, hash := range configHashesByTable {
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
