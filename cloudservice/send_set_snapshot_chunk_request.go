package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendSetSnapshotChunkRequest(snapshotId string, number uint32, data []byte) (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_SetSnapshotChunkRequest{
			SetSnapshotChunkRequest: &emailproto.SetSnapshotChunkRequest{
				SnapshotId: snapshotId,
				Number:     number,
				Data:       data,
			},
		},
	}
	return self.SendRequest(req)
}
