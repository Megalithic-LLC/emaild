package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/rs/xid"
)

func (self MessagesDAO) CreateTx(tx *genji.Tx, message *model.Message) error {
	if table, err := tx.GetTable(model.MessageTable); err != nil {
		return err
	} else {
		if message.ID == "" {
			message.ID = xid.New().String()
		}
		_, err := table.Insert(message)
		return err
	}
}

func (self MessagesDAO) FindByIDTx(tx *genji.Tx, id string) (*model.Message, error) {
	messageTable, err := tx.GetTable(model.MessageTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Message{ID: id}
	messagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := messageTable.GetRecord(messagePK)
	if err != nil {
		return nil, err
	}
	var message model.Message
	if err := message.ScanRecord(r); err != nil {
		return nil, err
	}
	return &message, nil
}
