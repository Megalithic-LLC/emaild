package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendSnapshotProgressRequest(snapshotId string, progress float32, size uint64) (*emailproto.ServerMessage, error) {
	snapshotProgressReq := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_SnapshotProgressRequest{
			SnapshotProgressRequest: &emailproto.SnapshotProgressRequest{
				SnapshotId: snapshotId,
				Progress:   progress,
				Size:       size,
			},
		},
	}
	return self.SendRequest(snapshotProgressReq)
}
