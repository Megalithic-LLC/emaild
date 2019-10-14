package snapshotmanager

import (
	"io/ioutil"
	"os"

	"github.com/docktermj/go-logger/logger"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

func (self *SnapshotManager) expungeSnapshotFiles(snapshots []*model.Snapshot) error {
	logger.Trace("SnapshotManager.expungeSnapshotFiles()")

	snapshotsDir, err := self.getSnapshotsDir()
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(snapshotsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		recordExistsForFile := false
		for _, snapshot := range snapshots {
			if file.Name() == self.getSnapshotFileName(snapshot) {
				recordExistsForFile = true
				break
			}
		}
		if !recordExistsForFile {
			filepath := snapshotsDir + "/" + file.Name()
			if err := os.Remove(filepath); err != nil {
				logger.Errorf("Failure expunging snapshot file %s: %v", file.Name(), err)
			} else {
				logger.Infof("Expunged snapshot file %s", file.Name())
			}
		}
	}

	return nil
}
