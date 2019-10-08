package dao

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/rs/xid"
)

func (self MailboxesDAO) AllocateNextUidTx(tx *genji.Tx, mailbox *model.Mailbox) (uint32, error) {
	uidNext := mailbox.UidNext
	mailbox.UidNext++
	if err := self.ReplaceTx(tx, mailbox); err != nil {
		mailbox.UidNext--
		return 0, err
	}
	return uidNext, nil
}

func (self MailboxesDAO) CreateTx(tx *genji.Tx, mailbox *model.Mailbox) error {
	if mailboxTable, err := tx.GetTable(model.MailboxTable); err != nil {
		return err
	} else {
		if mailbox.Id == "" {
			mailbox.Id = xid.New().String()
		}
		_, err := mailboxTable.Insert(mailbox)
		return err
	}
}

func (self MailboxesDAO) DeleteTx(tx *genji.Tx, id string) error {
	mailboxTable, err := tx.GetTable(model.MailboxTable)
	if err != nil {
		return err
	}
	searchFor := &model.Mailbox{Id: id}
	mailboxPK, err := searchFor.PrimaryKey()
	if err != nil {
		return err
	}
	return mailboxTable.Delete(mailboxPK)
}

func (self MailboxesDAO) FindTx(tx *genji.Tx, where query.Expr, limit int, iter func(mailbox *model.Mailbox) error) error {
	mailboxTable, err := tx.GetTable(model.MailboxTable)
	if err != nil {
		return err
	}
	selectStmt := query.Select().From(mailboxTable)
	if where != nil {
		selectStmt = selectStmt.Where(where)
	}
	if limit > 0 {
		selectStmt = selectStmt.Limit(limit)
	}
	return selectStmt.Run(tx).Iterate(func(recordId []byte, r record.Record) error {
		var mailbox model.Mailbox
		if err := mailbox.ScanRecord(r); err != nil {
			return err
		}
		return iter(&mailbox)
	})
}

func (self MailboxesDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Mailbox, error) {
	mailboxTable, err := tx.GetTable(model.MailboxTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Mailbox{Id: id}
	mailboxPK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := mailboxTable.GetRecord(mailboxPK)
	if err != nil {
		return nil, err
	}
	var mailbox model.Mailbox
	if err := mailbox.ScanRecord(r); err != nil {
		return nil, err
	}
	return &mailbox, err
}

func (self MailboxesDAO) FindOneByNameTx(tx *genji.Tx, accountId string, name string) (*model.Mailbox, error) {
	var retval *model.Mailbox
	mailboxTable, err := tx.GetTable(model.MailboxTable)
	if err != nil {
		return nil, err
	}
	err = query.
		Select().
		From(mailboxTable).
		Where(
			query.And(
				self.fields.AccountId.Eq(accountId),
				self.fields.Name.Eq(name),
			),
		).
		Limit(1).
		Run(tx).
		Iterate(func(recordId []byte, r record.Record) error {
			var mailbox model.Mailbox
			if err := mailbox.ScanRecord(r); err == nil {
				retval = &mailbox
			}
			return err
		})
	return retval, err
}

func (self MailboxesDAO) ReplaceTx(tx *genji.Tx, mailbox *model.Mailbox) error {
	mailboxTable, err := tx.GetTable(model.MailboxTable)
	if err != nil {
		return err
	}
	pk, err := mailbox.PrimaryKey()
	if err != nil {
		return err
	}
	return mailboxTable.Replace(pk, mailbox)
}
