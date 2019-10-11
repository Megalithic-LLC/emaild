package submissionendpoint

import (
	"fmt"
	"net"

	"github.com/asdine/genji/query"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/model"
	"github.com/on-prem-net/emaild/submissionendpoint/submissionbackend"
)

type SubmissionEndpoint struct {
	endpointsDAO      dao.EndpointsDAO
	listeners         []net.Listener
	servers           []*smtp.Server
	submissionBackend *submissionbackend.SubmissionBackend
}

func New(
	endpointsDAO dao.EndpointsDAO,
	submissionBackend *submissionbackend.SubmissionBackend,
) *SubmissionEndpoint {
	self := SubmissionEndpoint{
		endpointsDAO:      endpointsDAO,
		listeners:         []net.Listener{},
		submissionBackend: submissionBackend,
		servers:           []*smtp.Server{},
	}
	self.Start()
	return &self
}

func (self *SubmissionEndpoint) Shutdown() {
	for _, listener := range self.listeners {
		listener.Close()
	}
	for _, server := range self.servers {
		server.Close()
	}
	self.listeners = []net.Listener{}
	self.servers = []*smtp.Server{}
	logger.Infof("SMTP Submission endpoint shutdown")
}

func (self *SubmissionEndpoint) Start() {
	logger.Tracef("SubmissionEndpoint:Start()")

	endpointFields := model.NewEndpointFields()
	where := query.And(
		endpointFields.Enabled.Eq(true),
		endpointFields.Protocol.Eq("submission"),
		endpointFields.Type.Eq("tcp"),
	)
	if err := self.endpointsDAO.Find(where, 0, func(endpoint *model.Endpoint) error {

		server := smtp.NewServer(self.submissionBackend)
		server.Addr = fmt.Sprintf(":%d", endpoint.Port)
		server.AllowInsecureAuth = true

		go func() {
			var err error
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				logger.Errorf("Failed listening: %v", err)
				return
			}

			logger.Infof("Listening for SMTP Submission on %v", server.Addr)

			if err := server.Serve(listener); err != nil {
				logger.Errorf("Failed listening: %v", err)
			} else {
				self.servers = append(self.servers, server)
				self.listeners = append(self.listeners, listener)
			}
		}()

		return nil

	}); err != nil {
		logger.Errorf("Failed loading endpoints: %v")
		return
	}
}
