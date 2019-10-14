package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
)

type MessageRawBodiesDAO struct {
	db     *genji.DB
	fields *model.MessageRawBodyFields
}

func NewMessageRawBodiesDAO(db *genji.DB) MessageRawBodiesDAO {
	return MessageRawBodiesDAO{
		db:     db,
		fields: model.NewMessageRawBodyFields(),
	}
}

func (self MessageRawBodiesDAO) Create(messageRawBody *model.MessageRawBody) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, messageRawBody)
	})
}

func (self MessageRawBodiesDAO) Delete(id string) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.DeleteTx(tx, id)
	})
}

func (self MessageRawBodiesDAO) FindById(id string) (*model.MessageRawBody, error) {
	var retval *model.MessageRawBody
	err := self.db.View(func(tx *genji.Tx) error {
		messageRawBody, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = messageRawBody
		}
		return err
	})
	return retval, err
}
