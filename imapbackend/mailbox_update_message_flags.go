package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) UpdateMessagesFlags(uid bool, seqset *imap.SeqSet, op imap.FlagsOp, flags []string) error {
	logger.Tracef("Mailbox:UpdateMessagesFlags()")
	return errors.New("NIY")
}
