package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/asdine/genji/table"
	"github.com/rs/xid"
)

type AccountsDAO struct {
	db     *genji.DB
	fields *model.AccountFields
}

func NewAccountsDAO(db *genji.DB) AccountsDAO {
	return AccountsDAO{
		db:     db,
		fields: model.NewAccountFields(),
	}
}

func (self AccountsDAO) Create(account *model.Account) error {
	return self.db.Update(func(tx *genji.Tx) error {
		if table, err := tx.GetTable(model.AccountTable); err != nil {
			return err
		} else {
			if account.Id == "" {
				account.Id = xid.New().String()
			}
			_, err := table.Insert(account)
			return err
		}
	})
}

func (self AccountsDAO) FindOneByUsername(username string) (*model.Account, error) {
	var retval *model.Account
	err := self.db.View(func(tx *genji.Tx) error {
		accountTable, err := tx.GetTable(model.AccountTable)
		if err != nil {
			return err
		}
		return query.
			Select().
			From(accountTable).
			Where(self.fields.Username.Eq(username)).
			Limit(1).
			Run(tx).
			Iterate(func(recordId []byte, r record.Record) error {
				var account model.Account
				if err := account.ScanRecord(r); err == nil {
					retval = &account
				}
				return err
			})
	})
	if retval == nil {
		return nil, table.ErrRecordNotFound
	}
	return retval, err
}

func (self AccountsDAO) FindByID(id string) (*model.Account, error) {
	var retval *model.Account
	err := self.db.View(func(tx *genji.Tx) error {
		accountTable, err := tx.GetTable(model.AccountTable)
		if err != nil {
			return err
		}
		accountSelector := &model.Account{Id: id}
		accountID, err := accountSelector.PrimaryKey()
		if err != nil {
			return err
		}
		r, err := accountTable.GetRecord(accountID)
		if err != nil {
			return err
		}
		var account model.Account
		err = account.ScanRecord(r)
		if err == nil {
			retval = &account
		}
		return err
	})
	return retval, err
}
