package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Reset() {
	logger.Tracef("Session:Reset()")
}
