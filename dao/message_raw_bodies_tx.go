package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
)

func (self MessageRawBodiesDAO) CreateTx(tx *genji.Tx, messageRawBody *model.MessageRawBody) error {
	if table, err := tx.GetTable(model.MessageRawBodyTable); err != nil {
		return err
	} else {
		_, err := table.Insert(messageRawBody)
		return err
	}
}

func (self MessageRawBodiesDAO) FindByIdTx(tx *genji.Tx, id string) (*model.MessageRawBody, error) {
	messageRawBodyTable, err := tx.GetTable(model.MessageRawBodyTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.MessageRawBody{Id: id}
	messageRawBodyPK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := messageRawBodyTable.GetRecord(messageRawBodyPK)
	if err != nil {
		return nil, err
	}
	var messageRawBody model.MessageRawBody
	if err := messageRawBody.ScanRecord(r); err != nil {
		return nil, err
	}
	return &messageRawBody, nil
}
