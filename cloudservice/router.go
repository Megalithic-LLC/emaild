package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
)

func (self *CloudService) route(message agentstreamproto.ServerMessage) {
	logger.Tracef("AgentStream:route()")

	switch message.MessageType.(type) {

	case *agentstreamproto.ServerMessage_ClaimRequest:
		claimRequest := message.GetClaimRequest()
		self.handleClaimRequest(message.Id, *claimRequest)
	case *agentstreamproto.ServerMessage_ConfigChangedRequest:
		configChangedRequest := message.GetConfigChangedRequest()
		self.handleConfigChangedRequest(message.Id, *configChangedRequest)
	}
}
