package dao

import (
	"github.com/asdine/genji"
	"github.com/on-prem-net/emaild/model"
)

type ServiceInstancesDAO struct {
	db     *genji.DB
	fields *model.ServiceInstanceFields
}

func NewServiceInstancesDAO(db *genji.DB) ServiceInstancesDAO {
	return ServiceInstancesDAO{
		db:     db,
		fields: model.NewServiceInstanceFields(),
	}
}

func (self ServiceInstancesDAO) Create(serviceInstance *model.ServiceInstance) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, serviceInstance)
	})
}

func (self ServiceInstancesDAO) FindById(id string) (*model.ServiceInstance, error) {
	var retval *model.ServiceInstance
	err := self.db.View(func(tx *genji.Tx) error {
		serviceInstance, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = serviceInstance
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}

func (self ServiceInstancesDAO) Replace(serviceInstance *model.ServiceInstance) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, serviceInstance)
	})
}
