package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/asdine/genji/table"
	"github.com/docktermj/go-logger/logger"
	"github.com/rs/xid"
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

func (self MailboxesDAO) Create(mailbox *model.Mailbox) error {
	return self.db.Update(func(tx *genji.Tx) error {
		if table, err := tx.GetTable(model.MailboxTable); err != nil {
			return err
		} else {
			if mailbox.ID == "" {
				mailbox.ID = xid.New().String()
			}
			_, err := table.Insert(mailbox)
			return err
		}
	})
}

func (self MailboxesDAO) FindByID(id string) (*model.Mailbox, error) {
	var retval *model.Mailbox
	err := self.db.View(func(tx *genji.Tx) error {
		mailboxTable, err := tx.GetTable(model.MailboxTable)
		if err != nil {
			return err
		}
		searchFor := &model.Mailbox{ID: id}
		mailboxID, err := searchFor.PrimaryKey()
		if err != nil {
			return err
		}
		r, err := mailboxTable.GetRecord(mailboxID)
		if err != nil {
			return err
		}
		var mailbox model.Mailbox
		err = mailbox.ScanRecord(r)
		if err == nil {
			retval = &mailbox
		}
		return err
	})
	return retval, err
}

func (self MailboxesDAO) FindOneByName(accountID string, name string) (*model.Mailbox, error) {
	logger.Tracef("MailboxesDAO:FindOneByName(%s, %s)", accountID, name)
	var retval *model.Mailbox
	err := self.db.View(func(tx *genji.Tx) error {
		mailboxTable, err := tx.GetTable(model.MailboxTable)
		if err != nil {
			return err
		}
		return query.
			Select().
			From(mailboxTable).
			Where(
				query.And(
					self.fields.AccountID.Eq(accountID),
					self.fields.Name.Eq(name),
				),
			).
			Limit(1).
			Run(tx).
			Iterate(func(recordID []byte, r record.Record) error {
				var mailbox model.Mailbox
				if err := mailbox.ScanRecord(r); err == nil {
					retval = &mailbox
				}
				return err
			})
	})
	if retval == nil {
		return nil, table.ErrRecordNotFound
	}
	return retval, err
}
