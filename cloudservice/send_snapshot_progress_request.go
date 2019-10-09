package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) SendSetSnapshotProgressRequest(snapshotId string, progress float32, size uint64) (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_SetSnapshotProgressRequest{
			SetSnapshotProgressRequest: &emailproto.SetSnapshotProgressRequest{
				SnapshotId: snapshotId,
				Progress:   progress,
				Size:       size,
			},
		},
	}
	return self.SendRequest(req)
}
