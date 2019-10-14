package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

func (self *CloudService) SendUpdateSnapshotRequest(snapshot *model.Snapshot) (*emailproto.ServerMessage, error) {

	pbSnapshot := SnapshotToProtobuf(snapshot)

	req := emailproto.ClientMessage{
		MessageType: &emailproto.ClientMessage_UpdateSnapshotRequest{
			UpdateSnapshotRequest: &emailproto.UpdateSnapshotRequest{
				Snapshot: pbSnapshot,
			},
		},
	}

	return self.SendRequest(req)
}
