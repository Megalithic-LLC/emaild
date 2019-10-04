package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Mail(from string) error {
	logger.Tracef("SMTP:Session:Mail(%s)", from)
	return nil
}
