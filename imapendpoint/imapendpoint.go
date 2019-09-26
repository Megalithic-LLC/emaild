package imapendpoint

import (
	"net"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
)

type ImapEndpoint struct {
	listener net.Listener
	server   *server.Server
}

func New(imapBackend backend.Backend) *ImapEndpoint {

	self := ImapEndpoint{}

	self.server = server.New(imapBackend)
	self.server.Addr = ":8143"
	self.server.AllowInsecureAuth = true

	go func() {
		var err error
		self.listener, err = net.Listen("tcp", self.server.Addr)
		if err != nil {
			logger.Fatalf("Failed listening: %v", err)
		}

		logger.Infof("Listening for IMAP4rev1 on %v", self.server.Addr)

		if err := self.server.Serve(self.listener); err != nil {
			logger.Errorf("Failed listening: %v", err)
		}
	}()

	return &self
}

func (self *ImapEndpoint) Shutdown() {
	self.server.Close()
	logger.Infof("IMAP endpoint shutdown")
}
