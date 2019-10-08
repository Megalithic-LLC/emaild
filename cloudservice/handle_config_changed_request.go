package cloudservice

import (
	"fmt"

	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
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
	for tableName, hash := range configHashesByTable {
		hashAsHex := fmt.Sprintf("%x", hash)
		key := fmt.Sprintf(propertykey.HashByTablePattern, tableName)
		if value, err := self.propertiesDAO.Get(key); err != nil {
			logger.Errorf("Failed looking up table config hash: %v", err)
		} else {
			if hashAsHex == value {
				continue
			}

			logger.Debugf("Table %s has changes", tableName)

			if tableName == "accounts" {
				if res, err := self.SendGetAccountsRequest(); err != nil {
					logger.Errorf("Failed requesting account changes: %v", err)
				} else if getAccountsRes := res.GetGetAccountsResponse(); getAccountsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						for _, pbAccount := range getAccountsRes.Accounts {
							account := AccountFromProtobuf(pbAccount)
							err := self.accountsDAO.ReplaceTx(tx, &account)
							if err == table.ErrRecordNotFound {
								err = self.accountsDAO.CreateTx(tx, &account)
							}
							if err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						logger.Errorf("Failed updating accounts: %v", err)
					} else {
						logger.Infof("Updated %d accounts", len(getAccountsRes.Accounts))
						self.propertiesDAO.Set(key, hashAsHex)
					}
				}
			}

		}
	}
}
