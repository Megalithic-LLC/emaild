package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
	"github.com/on-prem-net/emaild/model"
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
