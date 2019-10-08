package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
	"github.com/on-prem-net/emaild/model"
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
		return self.CreateTx(tx, account)
	})
}

func (self AccountsDAO) FindOneByEmail(email string) (*model.Account, error) {
	var retval *model.Account
	err := self.db.View(func(tx *genji.Tx) error {
		account, err := self.FindOneByEmailTx(tx, email)
		if err == nil {
			retval = account
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

func (self AccountsDAO) FindById(id string) (*model.Account, error) {
	var retval *model.Account
	err := self.db.View(func(tx *genji.Tx) error {
		account, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = account
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}
