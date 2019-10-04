package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Rcpt(to string) error {
	logger.Tracef("Session:Rcpt(%s)", to)
	return nil
}
