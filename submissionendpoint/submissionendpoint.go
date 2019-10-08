package submissionendpoint

import (
	"net"

	"github.com/on-prem-net/emaild/submissionendpoint/submissionbackend"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

type SubmissionEndpoint struct {
	listener net.Listener
	server   *smtp.Server
}

func New(submissionBackend *submissionbackend.SubmissionBackend) *SubmissionEndpoint {

	self := SubmissionEndpoint{}

	self.server = smtp.NewServer(submissionBackend)
	self.server.Addr = ":8587"
	self.server.AllowInsecureAuth = true

	go func() {
		var err error
		self.listener, err = net.Listen("tcp", self.server.Addr)
		if err != nil {
			logger.Fatalf("Failed listening: %v", err)
		}

		logger.Infof("Listening for SMTP Submission on %v", self.server.Addr)

		if err := self.server.Serve(self.listener); err != nil {
			logger.Errorf("Failed listening: %v", err)
		}
	}()

	return &self
}

func (self *SubmissionEndpoint) Shutdown() {
	self.server.Close()
	logger.Infof("SMTP Submission endpoint shutdown")
}
