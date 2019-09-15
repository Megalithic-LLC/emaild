package dao

import (
	"github.com/asdine/genji"
)

type PropertiesDAO struct {
	db *genji.DB
}

func NewPropertiesDAO(db *genji.DB) PropertiesDAO {
	return PropertiesDAO{db: db}
}

func (self PropertiesDAO) Get(key string) (retval string, err error) {
	self.db.View(func(tx *genji.Tx) error {
		retval, err = self.GetTx(tx, key)
		return err
	})
	return
}

func (self PropertiesDAO) Set(key, value string) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.SetTx(tx, key, value)
	})
}

func (self PropertiesDAO) SetIfKeyNotExists(key, value string) (retval string, err error) {
	self.db.Update(func(tx *genji.Tx) error {
		retval, err = self.SetIfKeyNotExistsTx(tx, key, value)
		return err
	})
	return
}
