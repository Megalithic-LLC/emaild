package smtpendpoint

import (
	"net"

	"github.com/on-prem-net/emaild/smtpendpoint/smtpbackend"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

type SmtpEndpoint struct {
	listener net.Listener
	server   *smtp.Server
}

func New(smtpBackend *smtpbackend.SmtpBackend) *SmtpEndpoint {

	self := SmtpEndpoint{}

	self.server = smtp.NewServer(smtpBackend)
	self.server.Addr = ":8025"
	self.server.AllowInsecureAuth = true

	go func() {
		var err error
		self.listener, err = net.Listen("tcp", self.server.Addr)
		if err != nil {
			logger.Fatalf("Failed listening: %v", err)
		}

		logger.Infof("Listening for SMTP on %v", self.server.Addr)

		if err := self.server.Serve(self.listener); err != nil {
			logger.Errorf("Failed listening: %v", err)
		}
	}()

	return &self
}

func (self *SmtpEndpoint) Shutdown() {
	self.server.Close()
	logger.Infof("SMTP endpoint shutdown")
}
