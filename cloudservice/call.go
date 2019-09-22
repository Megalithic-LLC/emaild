package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/agentstreamproto"
)

type Call struct {
	Done  chan bool
	Error error
	Req   agentstreamproto.ClientMessage
	Res   *agentstreamproto.ServerMessage
}
