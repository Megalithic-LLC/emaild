package main

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/rs/xid"
)

func getOrGenerateNodeID() (string, error) {

	// If a unique id does not yet exist, create it now
	var nodeid string
	if err := db.Update(func(tx *genji.Tx) error {
		table, err := tx.GetTable(model.PropertyTable)
		if err != nil {
			return err
		}

		f := model.NewPropertyFields()

		if err := query.
			Select().
			From(table).
			Where(f.Name.Eq("nodeid")).
			Run(tx).
			Iterate(func(recordID []byte, r record.Record) error {
				var property model.Property
				if err := property.ScanRecord(r); err != nil {
					return err
				}
				nodeid = property.Value
				return nil
			}); err != nil {
			return err
		}

		if nodeid == "" {
			nodeid = xid.New().String()
			if _, err := table.Insert(&model.Property{
				ID:    nodeid,
				Name:  "nodeid",
				Value: nodeid,
			}); err != nil {
				logger.Fatalf("Failed assigning node id: %v", err)
				return err
			}
			logger.Debugf("Assigned a new node id")
		}

		return nil
	}); err != nil {
		return "", err
	}

	return nodeid, nil
}
