package cloudservice

import (
	"errors"
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
)

type writer struct {
	prefix string
}

func (self writer) Write(p []byte) (int, error) {
	logger.Debugf("(%s) %s", self.prefix, string(p))
	return len(p), nil
}

func (self *CloudService) handleTestAccountCredentialsRequest(requestId uint64, req agentstreamproto.EmailcdnTestAccountCredentialsRequest) {
	logger.Tracef("CloudService:handleTestAccountCredentialsRequest(%d, %+v)", requestId, req)

	if req.Account.Provider == "imap" {

		// Verify we can connect to host
		addr := fmt.Sprintf("%s:%d", req.Account.ImapHost, req.Account.ImapPort)
		var c *client.Client
		if req.Account.SslRequired {
			logger.Debugf("Connecting to %s via TLS", addr)
			var err error
			c, err = client.DialTLS(addr, nil)
			if err != nil {
				logger.Debugf("Dail TLS failed: %v", err)
				self.SendErrorResponse(requestId, err)
				return
			}
			logger.Infof("Connected to %s via TLS", addr)
		} else {
			logger.Debugf("Connecting to %s", addr)
			var err error
			c, err = client.Dial(addr)
			if err != nil {
				logger.Infof("Dail failed: %v", err)
				self.SendErrorResponse(requestId, err)
				return
			}
			logger.Infof("Connected to %s", addr)
		}
		defer c.Logout()

		recvWriter := writer{prefix: "recv"}
		sendWriter := writer{prefix: "send"}
		c.SetDebug(imap.NewDebugWriter(sendWriter, recvWriter))

		// Verify credentials with AUTH=PLAIN
		if ok, err := c.Support("AUTH=PLAIN"); err != nil {
			logger.Debugf("Failed detecting capabilities: %v", err)
			self.SendErrorResponse(requestId, err)
			return
		} else if ok {
			saslClient := sasl.NewPlainClient("", req.Account.ImapUsername, req.Account.ImapPassword)
			if err := c.Authenticate(saslClient); err != nil {
				msg := fmt.Sprintf("Authentication failed: %v", err)
				logger.Debugf(msg)
				self.SendErrorResponse(requestId, errors.New(msg))
				return
			}
			logger.Debugf("Authenticate succeeded")
			self.SendAckResponse(requestId)
		}

		self.SendErrorResponse(requestId, errors.New("Unsupported auth mechanism"))
	}

	self.SendErrorResponse(requestId, errors.New("Unsupported provider: "+req.Account.Provider))
}
