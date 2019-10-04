package smtpbackend

import (
	"io"

	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Data(r io.Reader) error {
	logger.Tracef("Session:Data()")
	return nil
}
