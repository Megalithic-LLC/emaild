package dao

import (
	"github.com/asdine/genji"
	"github.com/on-prem-net/emaild/model"
	"github.com/rs/xid"
)

func (self SnapshotsDAO) CreateTx(tx *genji.Tx, snapshot *model.Snapshot) error {
	if snapshotTable, err := tx.GetTable(model.SnapshotTable); err != nil {
		return err
	} else {
		if snapshot.Id == "" {
			snapshot.Id = xid.New().String()
		}
		_, err := snapshotTable.Insert(snapshot)
		return err
	}
}

func (self SnapshotsDAO) FindByIdTx(tx *genji.Tx, id string) (*model.Snapshot, error) {
	snapshotTable, err := tx.GetTable(model.SnapshotTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.Snapshot{Id: id}
	pk, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := snapshotTable.GetRecord(pk)
	if err != nil {
		return nil, err
	}
	var snapshot model.Snapshot
	err = snapshot.ScanRecord(r)
	return &snapshot, err
}

func (self SnapshotsDAO) ReplaceTx(tx *genji.Tx, snapshot *model.Snapshot) error {
	snapshotTable, err := tx.GetTable(model.SnapshotTable)
	if err != nil {
		return err
	}
	pk, err := snapshot.PrimaryKey()
	if err != nil {
		return err
	}
	return snapshotTable.Replace(pk, snapshot)
}
