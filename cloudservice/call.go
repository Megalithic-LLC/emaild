package cloudservice

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice/emailproto"
)

type Call struct {
	Done  chan bool
	Error error
	Req   emailproto.ClientMessage
	Res   *emailproto.ServerMessage
}
