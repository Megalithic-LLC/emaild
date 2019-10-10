package imapendpoint

import (
	"fmt"
	"net"

	"github.com/asdine/genji/query"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/model"
)

type ImapEndpoint struct {
	endpointsDAO dao.EndpointsDAO
	imapBackend  backend.Backend
	listeners    []net.Listener
	servers      []*server.Server
}

func New(
	endpointsDAO dao.EndpointsDAO,
	imapBackend backend.Backend,
) *ImapEndpoint {
	self := ImapEndpoint{
		endpointsDAO: endpointsDAO,
		imapBackend:  imapBackend,
		listeners:    []net.Listener{},
		servers:      []*server.Server{},
	}
	self.Start()
	return &self
}

func (self *ImapEndpoint) Shutdown() {
	for _, listener := range self.listeners {
		listener.Close()
	}
	for _, server := range self.servers {
		server.Close()
	}
	self.listeners = []net.Listener{}
	self.servers = []*server.Server{}
	logger.Infof("IMAP endpoint shutdown")
}

func (self *ImapEndpoint) Start() {
	logger.Tracef("ImapEndpoint:Start()")

	endpointFields := model.NewEndpointFields()
	where := query.And(
		endpointFields.Enabled.Eq(true),
		endpointFields.Protocol.Eq("imap"),
		endpointFields.Type.Eq("tcp"),
	)
	if err := self.endpointsDAO.Find(where, 0, func(endpoint *model.Endpoint) error {

		server := server.New(self.imapBackend)
		server.Addr = fmt.Sprintf(":%d", endpoint.Port)
		server.AllowInsecureAuth = true

		go func() {
			var err error
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				logger.Fatalf("Failed listening: %v", err)
			}

			logger.Infof("Listening for IMAP4rev1 on %v", server.Addr)

			if err := server.Serve(listener); err != nil {
				logger.Errorf("Failed listening: %v", err)
			} else {
				self.listeners = append(self.listeners, listener)
				self.servers = append(self.servers, server)
			}
		}()

		return nil
	}); err != nil {
		logger.Errorf("Failed loading endpoints: %v")
		return
	}
}
