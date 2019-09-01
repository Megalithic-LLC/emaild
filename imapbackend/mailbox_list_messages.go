package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) ListMessages(uid bool, seqSet *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	logger.Tracef("Mailbox:ListMessages()")
	return errors.New("NIY")
}
