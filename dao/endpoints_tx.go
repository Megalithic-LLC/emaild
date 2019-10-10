package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
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

func (self EndpointsDAO) FindTx(tx *genji.Tx, where query.Expr, limit int, iter func(endpoint *model.Endpoint) error) error {
	endpointTable, err := tx.GetTable(model.EndpointTable)
	if err != nil {
		return err
	}
	selectStmt := query.Select().From(endpointTable)
	if where != nil {
		selectStmt = selectStmt.Where(where)
	}
	if limit > 0 {
		selectStmt = selectStmt.Limit(limit)
	}
	return selectStmt.Run(tx).Iterate(func(recordId []byte, r record.Record) error {
		var endpoint model.Endpoint
		if err := endpoint.ScanRecord(r); err != nil {
			return err
		}
		return iter(&endpoint)
	})
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
