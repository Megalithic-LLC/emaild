package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
)

type MessagesDAO struct {
	db     *genji.DB
	fields *model.MessageFields
}

func NewMessagesDAO(db *genji.DB) MessagesDAO {
	return MessagesDAO{
		db:     db,
		fields: model.NewMessageFields(),
	}
}

func (self MessagesDAO) Create(message *model.Message) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, message)
	})
}

func (self MessagesDAO) FindById(id string) (*model.Message, error) {
	var retval *model.Message
	err := self.db.View(func(tx *genji.Tx) error {
		message, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = message
		}
		return err
	})
	return retval, err
}
