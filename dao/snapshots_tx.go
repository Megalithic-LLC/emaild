package dao

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
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

func (self SnapshotsDAO) FindAllTx(tx *genji.Tx) ([]*model.Snapshot, error) {
	snapshotTable, err := tx.GetTable(model.SnapshotTable)
	if err != nil {
		return nil, err
	}
	retval := []*model.Snapshot{}
	selectStmt := query.Select().From(snapshotTable)
	err = selectStmt.Run(tx).Iterate(func(recordId []byte, r record.Record) error {
		var snapshot model.Snapshot
		if err := snapshot.ScanRecord(r); err != nil {
			return err
		}
		retval = append(retval, &snapshot)
		return nil
	})
	return retval, err
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
