package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Mailbox) Check() error {
	logger.Tracef("Mailbox:Check()")
	return nil
}
