package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

func (self *CloudService) SendGetSnapshotChunksMissingRequest(snapshotId string) (*emailproto.ServerMessage, error) {
	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_GetSnapshotChunksMissingRequest{
			GetSnapshotChunksMissingRequest: &emailproto.GetSnapshotChunksMissingRequest{
				SnapshotId: snapshotId,
			},
		},
	}
	return self.SendRequest(req)
}
