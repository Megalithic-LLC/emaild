package cloudservice

import (
	"errors"
	"fmt"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/cloudservice/agentstreamproto"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
)

func (self *CloudService) handleTestAccountCredentialsRequest(requestId uint64, req agentstreamproto.EmailcdnTestAccountCredentialsRequest) {
	logger.Tracef("CloudService:handleTestAccountCredentialsRequest(%d, %+v)", requestId, req)

	if req.Account.Provider == "imap" {

		// Verify we can connect to host
		addr := fmt.Sprintf("%s:%d", req.Account.ImapHost, req.Account.ImapPort)
		var c *client.Client
		if req.Account.SslRequired {
			var err error
			c, err = client.DialTLS(addr, nil)
			if err != nil {
				logger.Debugf("Dail TLS failed: %v", err)
				self.SendErrorResponse(requestId, err)
				return
			}
			logger.Debugf("Connected to %s:%d via TLS", req.Account.ImapHost, req.Account.ImapPort)
		} else {
			var err error
			c, err = client.Dial(addr)
			if err != nil {
				logger.Debugf("Dail failed: %v", err)
				self.SendErrorResponse(requestId, err)
				return
			}
		}
		defer c.Logout()

		//c.SetDebug(imap.NewDebugWriter(os.Stderr, os.Stderr))

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
