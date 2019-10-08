package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

type Call struct {
	Done  chan bool
	Error error
	Req   emailproto.ClientMessage
	Res   *emailproto.ServerMessage
}
