package dao

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
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

func (self MailboxMessagesDAO) Delete(mailboxId, messageId string) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.DeleteTx(tx, mailboxId, messageId)
	})
}

func (self MailboxMessagesDAO) Find(where query.Expr, limit int, iter func(mailboxMessage *model.MailboxMessage) error) error {
	return self.db.View(func(tx *genji.Tx) error {
		return self.FindTx(tx, where, limit, iter)
	})
}

func (self MailboxMessagesDAO) FindByIds(mailboxId, messageId string) (*model.MailboxMessage, error) {
	var retval *model.MailboxMessage
	err := self.db.View(func(tx *genji.Tx) error {
		mailboxMessage, err := self.FindByIdsTx(tx, mailboxId, messageId)
		if err == nil {
			retval = mailboxMessage
		}
		return err
	})
	return retval, err
}

func (self MailboxMessagesDAO) FindForUpdate(where query.Expr, limit int, iter func(mailboxMessage *model.MailboxMessage) error) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.FindTx(tx, where, limit, iter)
	})
}

func (self MailboxMessagesDAO) Replace(mailboxMessage *model.MailboxMessage) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, mailboxMessage)
	})
}
