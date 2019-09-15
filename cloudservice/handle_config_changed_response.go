package cloudservice

import (
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/drauschenbach/megalithicd/propertykey"
)

func (self *CloudService) handleConfigChangedResponse(requestId uint64, configChangedRes agentstreamproto.ConfigChangedResponse) {
	logger.Tracef("CloudService:handleConfigChangedResponse(%d)", requestId)

	for table, hash := range configChangedRes.HashesByTable {
		hashAsHex := fmt.Sprintf("%x", hash)

		key := fmt.Sprintf(propertykey.HashByTablePattern, table)
		if value, err := self.propertiesDAO.Get(key); err != nil {
			logger.Errorf("Failed looking up table config hash: %v", err)
		} else {
			if hashAsHex == value {
				continue
			}

			if table == "emailcdnAccounts" {
				getAccountsRes, err := self.SendGetAccountsRequest()
				if err != nil {
					logger.Errorf("Failed requesting accounts: %v", err)
				} else {
					for _, pbAccount := range getAccountsRes.Accounts {
						account := &model.Account{
							ID:           pbAccount.Id,
							Name:         pbAccount.Name,
							Provider:     pbAccount.Provider,
							Email:        pbAccount.Email,
							ImapHost:     pbAccount.ImapHost,
							ImapPort:     uint16(pbAccount.ImapPort),
							ImapUsername: pbAccount.ImapUsername,
							ImapPassword: pbAccount.ImapPassword,
							SmtpHost:     pbAccount.SmtpHost,
							SmtpPort:     uint16(pbAccount.SmtpPort),
							SmtpUsername: pbAccount.SmtpUsername,
							SmtpPassword: pbAccount.SmtpPassword,
							SslRequired:  pbAccount.SslRequired,
						}

						if err := self.accountsDAO.Upsert(account); err != nil {
							logger.Errorf("Failed storing account: %v", err)
							return
						}
					}
				}
			}

			self.propertiesDAO.Set(key, hashAsHex)
		}
	}

}
