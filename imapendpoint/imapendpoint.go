package imapendpoint

import (
	"github.com/emersion/go-imap/server"
)

type ImapEndpoint struct {
	listener net.Listener
	server   *http.Server
}

func New() *ImapEndpoint {
	self := ImapEndpoint{}
	
	self.server := server.New(imapBackend)
	self.server.Addr = ":8143"
	self.server.AllowInsecureAuth = true

	go func() {
		var err error
		self.listener, err = net.Listen("tcp", s.Addr)
		if err != nil {
			logger.Fatalf("Failed listening: %v", err)
		}

		logger.Infof("Listening for IMAP4rev1 on %v", self.server.Addr)

		if err := self.server.Serve(listener); err != nil {
			logger.Errorf("Failed listening: %v", err)
		}
	}()

	return &self
}

func (self *RestEndpoint) Shutdown(ctx context.Context) {
	self.server.Shutdown(ctx)
	logger.Infof("Rest endpoint shutdown")
}
