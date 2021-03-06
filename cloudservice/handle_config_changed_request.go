package cloudservice

import (
	"fmt"

	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
	"github.com/docktermj/go-logger/logger"
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
	"github.com/Megalithic-LLC/on-prem-emaild/propertykey"
)

func (self *CloudService) handleConfigChangedRequest(requestId uint64, configChangedReq emailproto.ConfigChangedRequest) {
	logger.Tracef("CloudService:handleConfigChangedRequest(%d)", requestId)

	if err := self.SendAckResponse(requestId); err != nil {
		logger.Errorf("Failed sending ack response: %v", err)
	}

	go self.processConfigChanges(configChangedReq.HashesByTable)
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

			switch tableName {

			case "accounts":
				if res, err := self.SendGetAccountsRequest(); err != nil {
					logger.Errorf("Failed requesting latest accounts: %v", err)
				} else if getAccountsRes := res.GetGetAccountsResponse(); getAccountsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						if err := self.accountsDAO.DeleteAllTx(tx); err != nil {
							return err
						}
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

			case "domains":
				if res, err := self.SendGetDomainsRequest(); err != nil {
					logger.Errorf("Failed requesting latest domains: %v", err)
				} else if getDomainsRes := res.GetGetDomainsResponse(); getDomainsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						if err := self.domainsDAO.DeleteAllTx(tx); err != nil {
							return err
						}
						for _, pbDomain := range getDomainsRes.Domains {
							domain := DomainFromProtobuf(pbDomain)
							err := self.domainsDAO.ReplaceTx(tx, &domain)
							if err == table.ErrRecordNotFound {
								err = self.domainsDAO.CreateTx(tx, &domain)
							}
							if err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						logger.Errorf("Failed updating domains: %v", err)
					} else {
						logger.Infof("Updated %d domains", len(getDomainsRes.Domains))
						self.propertiesDAO.Set(key, hashAsHex)
					}
				}

			case "endpoints":
				if res, err := self.SendGetEndpointsRequest(); err != nil {
					logger.Errorf("Failed requesting latest endpoints: %v", err)
				} else if getEndpointsRes := res.GetGetEndpointsResponse(); getEndpointsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						if err := self.endpointsDAO.DeleteAllTx(tx); err != nil {
							return err
						}
						for _, pbEndpoint := range getEndpointsRes.Endpoints {
							endpoint := EndpointFromProtobuf(pbEndpoint)
							err := self.endpointsDAO.ReplaceTx(tx, &endpoint)
							if err == table.ErrRecordNotFound {
								err = self.endpointsDAO.CreateTx(tx, &endpoint)
							}
							if err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						logger.Errorf("Failed updating endpoints: %v", err)
					} else {
						logger.Infof("Updated %d endpoints", len(getEndpointsRes.Endpoints))
						self.propertiesDAO.Set(key, hashAsHex)

						self.imapEndpoint.Shutdown()
						self.smtpEndpoint.Shutdown()
						self.submissionEndpoint.Shutdown()

						self.imapEndpoint.Start()
						self.smtpEndpoint.Start()
						self.submissionEndpoint.Start()
					}
				}

			case "snapshots":
				if res, err := self.SendGetSnapshotsRequest(); err != nil {
					logger.Errorf("Failed requesting latest shapshots: %v", err)
				} else if getSnapshotsRes := res.GetGetSnapshotsResponse(); getSnapshotsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						if err := self.snapshotsDAO.DeleteAllTx(tx); err != nil {
							return err
						}
						for _, pbSnapshot := range getSnapshotsRes.Snapshots {
							snapshot := SnapshotFromProtobuf(pbSnapshot)
							err := self.snapshotsDAO.ReplaceTx(tx, &snapshot)
							if err == table.ErrRecordNotFound {
								err = self.snapshotsDAO.CreateTx(tx, &snapshot)
							}
							if err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						logger.Errorf("Failed updating snapshots: %v", err)
					} else {
						logger.Infof("Updated %d snapshots", len(getSnapshotsRes.Snapshots))
						self.propertiesDAO.Set(key, hashAsHex)

						self.snapshotManager.Perform()
					}
				}

			}

		}
	}
}
