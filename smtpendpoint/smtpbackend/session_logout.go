package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Logout() error {
	logger.Tracef("Session:Logout()")
	return nil
}
