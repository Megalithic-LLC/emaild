package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type EndpointsDAO struct {
	db     *genji.DB
	fields *model.EndpointFields
}

func NewEndpointsDAO(db *genji.DB) EndpointsDAO {
	return EndpointsDAO{
		db:     db,
		fields: model.NewEndpointFields(),
	}
}

func (self EndpointsDAO) Create(endpoint *model.Endpoint) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, endpoint)
	})
}

func (self EndpointsDAO) Find(where query.Expr, limit int, iter func(endpoint *model.Endpoint) error) error {
	return self.db.View(func(tx *genji.Tx) error {
		return self.FindTx(tx, where, limit, iter)
	})
}

func (self EndpointsDAO) FindById(id string) (*model.Endpoint, error) {
	var retval *model.Endpoint
	err := self.db.View(func(tx *genji.Tx) error {
		endpoint, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = endpoint
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}

func (self EndpointsDAO) Replace(endpoint *model.Endpoint) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, endpoint)
	})
}
