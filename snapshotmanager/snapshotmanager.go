package snapshotmanager

import (
	"os"
	"sync"

	"github.com/asdine/genji/engine"
	"github.com/docktermj/go-logger/logger"
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/model"
)

type SnapshotManager struct {
	genjiEngine  engine.Engine
	listeners    []Listener
	mutex        sync.Mutex
	snapshotsDAO dao.SnapshotsDAO
}

type Listener interface {
	SnapshotProgress(snapshot *model.Snapshot, progress float32, size uint64)
}

func New(
	genjiEngine engine.Engine,
	snapshotsDAO dao.SnapshotsDAO,
) *SnapshotManager {
	self := SnapshotManager{
		genjiEngine:  genjiEngine,
		listeners:    []Listener{},
		snapshotsDAO: snapshotsDAO,
	}
	go self.Perform()
	return &self
}

func (self *SnapshotManager) RegisterListener(listener Listener) {
	self.listeners = append(self.listeners, listener)
}

func (self *SnapshotManager) NotifyListenersOfProgress(snapshot *model.Snapshot, progress float32, size uint64) {
	for _, listener := range self.listeners {
		go listener.SnapshotProgress(snapshot, progress, size)
	}
}

func (self *SnapshotManager) getSnapshotsDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		logger.Errorf("Unable to get working dir: %v", err)
		return "", err
	}
	snapshotsDir := dir + "/snapshots"
	if _, err := os.Stat(snapshotsDir); os.IsNotExist(err) {
		if err := os.Mkdir(snapshotsDir, os.ModePerm); err != nil {
			logger.Errorf("Failure creating snapshots dir: %v", err)
			return "", err
		}
	} else if err != nil {
		logger.Errorf("Unable to get snapshots dir: %v", err)
		return "", err
	}
	return snapshotsDir, nil
}

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

func (self *SnapshotManager) ensureSnapshotExists(snapshot *model.Snapshot) (os.FileInfo, error) {
	logger.Tracef("SnapshotManager.performSnapshot(%v)", snapshot)

	snapshotsDir, err := self.getSnapshotsDir()
	if err != nil {
		return nil, err
	}

	filename := "snapshot-" + snapshot.Id + ".db"
	filepath := snapshotsDir + "/" + filename

	if fileInfo, err := os.Stat(filepath); err == nil {
		logger.Debugf("Snapshot %s already exists at %s", snapshot.Id, filepath)
		return fileInfo, nil
	} else if !os.IsNotExist(err) {
		logger.Errorf("Failure checking for presence of snapshot: %v", err)
		return nil, err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	if err := self.performSnapshot(snapshot, file); err != nil {
		os.Remove(filepath)
		return nil, err
	}

	return file.Stat()
}

func (self *SnapshotManager) getSnapshotFileName(snapshot *model.Snapshot) string {
	return "snapshot-" + snapshot.Id + ".db"
}
