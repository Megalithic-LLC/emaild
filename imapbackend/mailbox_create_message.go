package imapbackend

import (
	"errors"
	"time"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	logger.Tracef("Mailbox:CreateMessage()")
	return errors.New("NIY")
}
