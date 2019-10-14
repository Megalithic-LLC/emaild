package snapshotmanager

import (
	"fmt"
	"os"

	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

// Perform any work to get things caught up with desired state
func (self *SnapshotManager) Perform() {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	logger.Trace("SnapshotManager.Perform()")
	snapshots, err := self.snapshotsDAO.FindAll()
	if err != nil {
		logger.Errorf("Failed loading snapshots: %v")
		return
	}
	logger.Debugf("Got %d snapshots", len(snapshots))

	// Create snapshots as needed
	for _, snapshot := range snapshots {
		self.ensureSnapshotExists(snapshot)
	}

	// Expunge snapshots as needed
	self.expungeSnapshotFiles(snapshots)
}

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

		snapshot.Engine = "bolt"
		snapshot.Progress = 50
		snapshot.Size = uint64(n)

		self.NotifyListeners(snapshot)

		return nil
	} else {
		return fmt.Errorf("Snapshotting database engine %T not supported yet", self.genjiEngine)
	}
}
