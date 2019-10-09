package dao

import (
	"github.com/asdine/genji"
	"github.com/on-prem-net/emaild/model"
)

type SnapshotsDAO struct {
	db     *genji.DB
	fields *model.SnapshotFields
}

func NewSnapshotsDAO(db *genji.DB) SnapshotsDAO {
	return SnapshotsDAO{
		db:     db,
		fields: model.NewSnapshotFields(),
	}
}

func (self SnapshotsDAO) Create(snapshot *model.Snapshot) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.CreateTx(tx, snapshot)
	})
}

func (self SnapshotsDAO) FindAll() ([]*model.Snapshot, error) {
	retval := []*model.Snapshot{}
	err := self.db.Update(func(tx *genji.Tx) error {
		snapshots, err := self.FindAllTx(tx)
		if err == nil {
			retval = snapshots
		}
		return err
	})
	return retval, err
}

func (self SnapshotsDAO) FindById(id string) (*model.Snapshot, error) {
	var retval *model.Snapshot
	err := self.db.View(func(tx *genji.Tx) error {
		snapshot, err := self.FindByIdTx(tx, id)
		if err == nil {
			retval = snapshot
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}

func (self SnapshotsDAO) Replace(snapshot *model.Snapshot) error {
	return self.db.Update(func(tx *genji.Tx) error {
		return self.ReplaceTx(tx, snapshot)
	})
}
