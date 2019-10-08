package cloudservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/on-prem-net/emaild/cloudservice/emailproto"
)

func (self *CloudService) route(message emailproto.ServerMessage) {
	logger.Tracef("AgentStream:route()")

	switch message.MessageType.(type) {

	case *emailproto.ServerMessage_ClaimRequest:
		claimRequest := message.GetClaimRequest()
		self.handleClaimRequest(message.Id, *claimRequest)

	case *emailproto.ServerMessage_ConfigChangedRequest:
		configChangedRequest := message.GetConfigChangedRequest()
		self.handleConfigChangedRequest(message.Id, *configChangedRequest)
	}
}
