package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

func (self *SmtpBackend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	logger.Tracef("SmtpBackend:AnonymousLogin()")
	return nil, smtp.ErrAuthRequired
}
