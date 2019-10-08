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

			switch tableName {

			case "accounts":
				if res, err := self.SendGetAccountsRequest(); err != nil {
					logger.Errorf("Failed requesting latest accounts: %v", err)
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

			case "serviceInstances":
				if res, err := self.SendGetServiceInstancesRequest(); err != nil {
					logger.Errorf("Failed requesting latest service instances: %v", err)
				} else if getServiceInstancesRes := res.GetGetServiceInstancesResponse(); getServiceInstancesRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
						for _, pbServiceInstance := range getServiceInstancesRes.ServiceInstances {
							serviceInstance := ServiceInstanceFromProtobuf(pbServiceInstance)
							err := self.serviceInstancesDAO.ReplaceTx(tx, &serviceInstance)
							if err == table.ErrRecordNotFound {
								err = self.serviceInstancesDAO.CreateTx(tx, &serviceInstance)
							}
							if err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						logger.Errorf("Failed updating service instances: %v", err)
					} else {
						logger.Infof("Updated %d service instances", len(getServiceInstancesRes.ServiceInstances))
						self.propertiesDAO.Set(key, hashAsHex)
					}
				}

			case "snapshots":
				if res, err := self.SendGetSnapshotsRequest(); err != nil {
					logger.Errorf("Failed requesting latest shapshots: %v", err)
				} else if getSnapshotsRes := res.GetGetSnapshotsResponse(); getSnapshotsRes != nil {
					if err := self.db.Update(func(tx *genji.Tx) error {
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
					}
				}

			}

		}
	}
}
