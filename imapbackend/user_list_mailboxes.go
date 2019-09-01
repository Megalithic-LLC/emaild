package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
)

func (self *User) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	logger.Tracef("User:ListMailboxes()")
	return nil, errors.New("NIY")
}
