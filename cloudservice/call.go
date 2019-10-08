package cloudservice

import (
	"github.com/on-prem-net/emaild/cloudservice/agentstreamproto"
)

type Call struct {
	Done  chan bool
	Error error
	Req   agentstreamproto.ClientMessage
	Res   *agentstreamproto.ServerMessage
}
