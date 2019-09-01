package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
)

func (self *Mailbox) Expunge() error {
	logger.Tracef("Mailbox:Expunge()")
	return errors.New("NIY")
}
