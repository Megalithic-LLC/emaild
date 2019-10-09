package snapshotmanager

import (
	"fmt"
	"os"

	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
	"github.com/on-prem-net/emaild/model"
)

func (self *SnapshotManager) performSnapshot(snapshot *model.Snapshot, file *os.File) error {
	logger.Tracef("SnapshotManager.performSnapshot(%s, %s)", snapshot.Id, file.Name())

	if boltEngine, ok := self.genjiEngine.(*bolt.Engine); ok {
		db := boltEngine.DB
		tx, err := db.Begin(false)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		n, err := tx.WriteTo(file)
		if err != nil {
			return err
		}
		logger.Infof("Wrote %d bytes to %s", n, file.Name())
		if err := file.Close(); err != nil {
			logger.Errorf("Failure closing snapshot file: %v", err)
			return err
		}
		self.NotifyListenersOfProgress(snapshot, 50.0, uint64(n))
		return nil
	} else {
		return fmt.Errorf("Snapshotting database engine %T not supported yet", self.genjiEngine)
	}
}
