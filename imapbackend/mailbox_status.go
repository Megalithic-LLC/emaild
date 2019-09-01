package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) Status(items []imap.StatusItem) (*imap.MailboxStatus, error) {
	logger.Tracef("Mailbox:Status()")
	return nil, errors.New("NIY")
}
