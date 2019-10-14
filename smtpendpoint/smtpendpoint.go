package smtpendpoint

import (
	"fmt"
	"net"

	"github.com/asdine/genji/query"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/Megalithic-LLC/on-prem-emaild/smtpendpoint/smtpbackend"
)

type SmtpEndpoint struct {
	endpointsDAO dao.EndpointsDAO
	listeners    []net.Listener
	servers      []*smtp.Server
	smtpBackend  *smtpbackend.SmtpBackend
}

func New(
	endpointsDAO dao.EndpointsDAO,
	smtpBackend *smtpbackend.SmtpBackend,
) *SmtpEndpoint {
	self := SmtpEndpoint{
		endpointsDAO: endpointsDAO,
		listeners:    []net.Listener{},
		servers:      []*smtp.Server{},
		smtpBackend:  smtpBackend,
	}
	self.Start()
	return &self
}

func (self *SmtpEndpoint) Shutdown() {
	for _, listener := range self.listeners {
		listener.Close()
	}
	for _, server := range self.servers {
		server.Close()
	}
	self.listeners = []net.Listener{}
	self.servers = []*smtp.Server{}
	logger.Infof("SMTP endpoint shutdown")
}

func (self *SmtpEndpoint) Start() {
	logger.Tracef("SmtpEndpoint:Start()")

	endpointFields := model.NewEndpointFields()
	where := query.And(
		endpointFields.Enabled.Eq(true),
		endpointFields.Protocol.Eq("smtp"),
		endpointFields.Type.Eq("tcp"),
	)
	if err := self.endpointsDAO.Find(where, 0, func(endpoint *model.Endpoint) error {

		server := smtp.NewServer(self.smtpBackend)
		server.Addr = fmt.Sprintf(":%d", endpoint.Port)
		server.AllowInsecureAuth = true

		go func() {
			var err error
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				logger.Errorf("Failed listening: %v", err)
				return
			}

			logger.Infof("Listening for SMTP on %v", server.Addr)

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
