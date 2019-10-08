package smtpbackend

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

func (self *SmtpBackend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	logger.Tracef("SmtpBackend:AnonymousLogin()")
	session := Session{
		backend:    self,
		recipients: []*model.Account{},
	}
	return &session, nil
}
