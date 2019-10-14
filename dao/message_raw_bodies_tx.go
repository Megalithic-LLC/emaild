package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
)

func (self MessageRawBodiesDAO) CreateTx(tx *genji.Tx, messageRawBody *model.MessageRawBody) error {
	if messageRawBodyTable, err := tx.GetTable(model.MessageRawBodyTable); err != nil {
		return err
	} else {
		_, err := messageRawBodyTable.Insert(messageRawBody)
		return err
	}
}

func (self MessageRawBodiesDAO) DeleteTx(tx *genji.Tx, messageId string) error {
	messageRawBodyTable, err := tx.GetTable(model.MessageRawBodyTable)
	if err != nil {
		return err
	}
	searchFor := &model.MessageRawBody{Id: messageId}
	messagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return err
	}
	return messageRawBodyTable.Delete(messagePK)
}

func (self MessageRawBodiesDAO) FindByIdTx(tx *genji.Tx, id string) (*model.MessageRawBody, error) {
	messageRawBodyTable, err := tx.GetTable(model.MessageRawBodyTable)
	if err != nil {
		if err == table.ErrRecordNotFound {
			return nil, nil
		}
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
