package dao

import (
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
)

type AccountsDAO struct {
	db *genji.DB
}

func NewAccountsDAO(db *genji.DB) AccountsDAO {
	return AccountsDAO{db: db}
}

func (self AccountsDAO) Create(account *model.Account) (err error) {
	self.db.Update(func(tx *genji.Tx) error {
		err = self.CreateTx(tx, account)
		return err
	})
	if err == nil {
		logger.Debug("Created account")
	}
	return
}

func (self AccountsDAO) DeleteByID(id string) (err error) {
	self.db.Update(func(tx *genji.Tx) error {
		err = self.DeleteByIDTx(tx, id)
		return err
	})
	if err == nil {
		logger.Debug("Deleted account")
	}
	return
}

func (self AccountsDAO) FindByID(id string) (retval *model.Account, err error) {
	self.db.View(func(tx *genji.Tx) error {
		retval, err = self.FindByIDTx(tx, id)
		return err
	})
	return
}

func (self AccountsDAO) Upsert(account *model.Account) (err error) {
	self.db.Update(func(tx *genji.Tx) error {
		err = self.UpsertTx(tx, account)
		return err
	})
	if err == nil {
		logger.Debug("Updated account")
	}
	return
}
