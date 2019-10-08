package dao

import (
	"github.com/asdine/genji"
	"github.com/on-prem-net/emaild/model"
	"github.com/rs/xid"
)

func (self ServiceInstancesDAO) CreateTx(tx *genji.Tx, serviceInstance *model.ServiceInstance) error {
	if serviceInstanceTable, err := tx.GetTable(model.ServiceInstanceTable); err != nil {
		return err
	} else {
		if serviceInstance.Id == "" {
			serviceInstance.Id = xid.New().String()
		}
		_, err := serviceInstanceTable.Insert(serviceInstance)
		return err
	}
}

func (self ServiceInstancesDAO) FindByIdTx(tx *genji.Tx, id string) (*model.ServiceInstance, error) {
	serviceInstanceTable, err := tx.GetTable(model.ServiceInstanceTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.ServiceInstance{Id: id}
	pk, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := serviceInstanceTable.GetRecord(pk)
	if err != nil {
		return nil, err
	}
	var serviceInstance model.ServiceInstance
	err = serviceInstance.ScanRecord(r)
	return &serviceInstance, err
}

func (self ServiceInstancesDAO) ReplaceTx(tx *genji.Tx, serviceInstance *model.ServiceInstance) error {
	serviceInstanceTable, err := tx.GetTable(model.ServiceInstanceTable)
	if err != nil {
		return err
	}
	pk, err := serviceInstance.PrimaryKey()
	if err != nil {
		return err
	}
	return serviceInstanceTable.Replace(pk, serviceInstance)
}
