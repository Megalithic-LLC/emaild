package cloudservice

import (
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
)

type Call struct {
	Done  chan bool
	Error error
	Req   agentstreamproto.ClientMessage
	Res   *agentstreamproto.ServerMessage
}
