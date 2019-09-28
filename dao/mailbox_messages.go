package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

type MailboxMessagesDAO struct {
	db     *genji.DB
	fields *model.MailboxMessageFields
}

func NewMailboxMessagesDAO(db *genji.DB) MailboxMessagesDAO {
	return MailboxMessagesDAO{
		db:     db,
		fields: model.NewMailboxMessageFields(),
	}
}

func (self MailboxMessagesDAO) Create(mailboxMessage *model.MailboxMessage) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, mailboxMessage)
	})
}

func (self MailboxMessagesDAO) Find(where query.Expr, limit int, iter func(recordID []byte, r record.Record) error) error {
	return self.db.View(func(tx *genji.Tx) error {
		return self.FindTx(tx, where, limit, iter)
	})
}

func (self MailboxMessagesDAO) FindByIDs(mailboxID, messageID string) (*model.MailboxMessage, error) {
	var retval *model.MailboxMessage
	err := self.db.View(func(tx *genji.Tx) error {
		mailboxMessage, err := self.FindByIDsTx(tx, mailboxID, messageID)
		if err == nil {
			retval = mailboxMessage
		}
		return err
	})
	return retval, err
}
