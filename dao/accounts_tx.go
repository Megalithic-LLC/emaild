package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/drauschenbach/megalithicd/model"
)

func (self AccountsDAO) CreateTx(tx *genji.Tx, account *model.Account) error {
	if table, err := tx.GetTable(model.AccountTable); err != nil {
		return err
	} else {
		_, err := table.Insert(account)
		return err
	}
}

func (self AccountsDAO) DeleteByIDTx(tx *genji.Tx, id string) error {
	accountTable, err := tx.GetTable(model.AccountTable)
	if err != nil {
		return err
	}
	f := model.NewAccountFields()
	return query.Delete().From(accountTable).Where(f.ID.Eq(id)).Run(tx)
}

func (self AccountsDAO) FindByIDTx(tx *genji.Tx, id string) (*model.Account, error) {
	accountTable, err := tx.GetTable(model.AccountTable)
	if err != nil {
		return nil, err
	}
	f := model.NewAccountFields()
	var retval model.Account
	err = query.
		Select().
		From(accountTable).
		Where(f.ID.Eq(id)).
		Limit(1).
		Run(tx).
		Iterate(func(recordID []byte, r record.Record) error {
			return retval.ScanRecord(r)
		})
	return &retval, err
}

func (self AccountsDAO) UpsertTx(tx *genji.Tx, account *model.Account) error {
	if err := self.DeleteByIDTx(tx, account.ID); err != nil {
		return err
	}
	return self.CreateTx(tx, account)
}
