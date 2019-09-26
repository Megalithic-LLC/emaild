package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) CopyMessages(uid bool, seqset *imap.SeqSet, destName string) error {
	logger.Tracef("Mailbox:CopyMessages()")
	return errors.New("NIY")
}
