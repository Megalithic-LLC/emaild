package dao

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
	"github.com/rs/xid"
)

func (self MessagesDAO) CreateTx(tx *genji.Tx, message *model.Message) error {
	if messageTable, err := tx.GetTable(model.MessageTable); err != nil {
		return err
	} else {
		if message.Id == "" {
			message.Id = xid.New().String()
		}
		_, err := messageTable.Insert(message)
		return err
	}
}

func (self MessagesDAO) DeleteTx(tx *genji.Tx, id string) error {
	messageTable, err := tx.GetTable(model.MessageTable)
	if err != nil {
		return err
	}
	searchFor := &model.Message{Id: id}
	messagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return err
	}
	return messageTable.Delete(messagePK)
}

func (self MessagesDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Message, error) {
	messageTable, err := tx.GetTable(model.MessageTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Message{Id: id}
	messagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := messageTable.GetRecord(messagePK)
	if err != nil {
		if err == table.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	var message model.Message
	if err := message.ScanRecord(r); err != nil {
		return nil, err
	}
	return &message, nil
}
