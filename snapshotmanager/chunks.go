package snapshotmanager

import (
	"os"

	"github.com/docktermj/go-logger/logger"
	"github.com/on-prem-net/emaild/model"
)

const (
	chunkSize = uint32(1000000)
)

func (self *SnapshotManager) GetChunk(snapshot *model.Snapshot, chunkNumber uint32) ([]byte, error) {
	logger.Tracef("SnapshotManager:GetChunk(%s, %d)", snapshot.Id, chunkNumber)

	snapshotsDir, err := self.getSnapshotsDir()
	if err != nil {
		return nil, err
	}

	filename := self.getSnapshotFileName(snapshot)
	filepath := snapshotsDir + "/" + filename

	file, err := os.Open(filepath)
	if err != nil {
		logger.Errorf("Failure reading snapshot file %s: %v", filepath, err)
		return nil, err
	}

	// get file size
	fileInfo, err := file.Stat()
	if err != nil {
		logger.Errorf("Failure reading snapshot file %s: %v", filepath, err)
		return nil, err
	}

	offset := chunkNumber * chunkSize
	bytesToRead := int64(chunkSize)
	if bytesRemaining := fileInfo.Size() - int64(offset); bytesRemaining < bytesToRead {
		bytesToRead = bytesRemaining
	}
	chunk := make([]byte, bytesToRead)
	if _, err := file.ReadAt(chunk, int64(offset)); err != nil {
		logger.Errorf("Failure reading snapshot file %s: %v", filepath, err)
		return nil, err
	}

	return chunk, nil
}
