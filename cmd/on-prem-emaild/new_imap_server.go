package main

import (
	"net"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
)

func newImapServer(imapBackend backend.Backend) *server.Server {
	s := server.New(imapBackend)
	s.Addr = ":8143"
	s.AllowInsecureAuth = true

	go func() {
		listener, err := net.Listen("tcp", s.Addr)
		if err != nil {
			logger.Fatalf("Failed listening: %v", err)
		}

		logger.Infof("Listening for IMAP4rev1 on %v", s.Addr)

		if err := s.Serve(listener); err != nil {
			logger.Errorf("Failed listening: %v", err)
		}
	}()

	return s
}
