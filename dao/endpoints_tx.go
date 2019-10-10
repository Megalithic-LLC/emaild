package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/on-prem-net/emaild/model"
	"github.com/rs/xid"
)

func (self EndpointsDAO) CreateTx(tx *genji.Tx, endpoint *model.Endpoint) error {
	if endpointTable, err := tx.GetTable(model.EndpointTable); err != nil {
		return err
	} else {
		if endpoint.Id == "" {
			endpoint.Id = xid.New().String()
		}
		_, err := endpointTable.Insert(endpoint)
		return err
	}
}

func (self EndpointsDAO) DeleteAllTx(tx *genji.Tx) error {
	endpointTable, err := tx.GetTable(model.EndpointTable)
	if err != nil {
		return err
	}
	return query.
		Delete().
		From(endpointTable).
		Run(tx)
}

func (self EndpointsDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Endpoint, error) {
	endpointTable, err := tx.GetTable(model.EndpointTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Endpoint{Id: id}
	pk, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := endpointTable.GetRecord(pk)
	if err != nil {
		return nil, err
	}
	var endpoint model.Endpoint
	err = endpoint.ScanRecord(r)
	return &endpoint, err
}

func (self EndpointsDAO) ReplaceTx(tx *genji.Tx, endpoint *model.Endpoint) error {
	endpointTable, err := tx.GetTable(model.EndpointTable)
	if err != nil {
		return err
	}
	pk, err := endpoint.PrimaryKey()
	if err != nil {
		return err
	}
	return endpointTable.Replace(pk, endpoint)
}
