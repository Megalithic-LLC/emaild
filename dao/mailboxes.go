package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/docktermj/go-logger/logger"
)

type MailboxesDAO struct {
	db     *genji.DB
	fields *model.MailboxFields
}

func NewMailboxesDAO(db *genji.DB) MailboxesDAO {
	return MailboxesDAO{
		db:     db,
		fields: model.NewMailboxFields(),
	}
}

func (self MailboxesDAO) AllocateNextUid(mailbox *model.Mailbox) (uint32, error) {
	var retval uint32
	err := self.db.Update(func(tx *genji.Tx) error {
		nextUid, err := self.AllocateNextUidTx(tx, mailbox)
		retval = nextUid
		return err
	})
	return retval, err
}

func (self MailboxesDAO) Create(mailbox *model.Mailbox) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, mailbox)
	})
}

func (self MailboxesDAO) Delete(id string) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.DeleteTx(tx, id)
	})
}

func (self MailboxesDAO) Find(where query.Expr, limit int, iter func(mailbox *model.Mailbox) error) error {
	return self.db.View(func(tx *genji.Tx) error {
		return self.FindTx(tx, where, limit, iter)
	})
}

func (self MailboxesDAO) FindById(id string) (*model.Mailbox, error) {
	var retval *model.Mailbox
	err := self.db.View(func(tx *genji.Tx) error {
		mailbox, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = mailbox
		}
		return err
	})
	return retval, err
}

func (self MailboxesDAO) FindOneByName(accountId string, name string) (*model.Mailbox, error) {
	logger.Tracef("MailboxesDAO:FindOneByName(%s, %s)", accountId, name)
	var retval *model.Mailbox
	err := self.db.View(func(tx *genji.Tx) error {
		mailbox, err := self.FindOneByNameTx(tx, accountId, name)
		if err == nil {
			retval = mailbox
		}
		return err
	})
	return retval, err
}

func (self MailboxesDAO) Replace(mailbox *model.Mailbox) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, mailbox)
	})
}
