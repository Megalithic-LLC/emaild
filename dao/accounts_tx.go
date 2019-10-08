package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/asdine/genji/table"
	"github.com/on-prem-net/emaild/model"
	"github.com/rs/xid"
)

func (self AccountsDAO) CreateTx(tx *genji.Tx, account *model.Account) error {
	if accountTable, err := tx.GetTable(model.AccountTable); err != nil {
		return err
	} else {
		if account.Id == "" {
			account.Id = xid.New().String()
		}
		_, err := accountTable.Insert(account)
		return err
	}
}

func (self AccountsDAO) FindOneByEmailTx(tx *genji.Tx, email string) (*model.Account, error) {
	var retval *model.Account
	accountTable, err := tx.GetTable(model.AccountTable)
	if err != nil {
		return nil, err
	}
	err = query.
		Select().
		From(accountTable).
		Where(self.fields.Email.Eq(email)).
		Limit(1).
		Run(tx).
		Iterate(func(recordId []byte, r record.Record) error {
			var account model.Account
			if err := account.ScanRecord(r); err == nil {
				retval = &account
			}
			return err
		})
	if err != nil {
		return nil, err
	}
	if retval == nil {
		return nil, table.ErrRecordNotFound
	}
	return retval, nil
}

func (self AccountsDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Account, error) {
	accountTable, err := tx.GetTable(model.AccountTable)
	if err != nil {
		return nil, err
	}
	accountSelector := &model.Account{Id: id}
	accountID, err := accountSelector.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := accountTable.GetRecord(accountID)
	if err != nil {
		return nil, err
	}
	var account model.Account
	err = account.ScanRecord(r)
	return &account, err
}

func (self AccountsDAO) ReplaceTx(tx *genji.Tx, account *model.Account) error {
	accountTable, err := tx.GetTable(model.AccountTable)
	if err != nil {
		return err
	}
	pk, err := account.PrimaryKey()
	if err != nil {
		return err
	}
	return accountTable.Replace(pk, account)
}
