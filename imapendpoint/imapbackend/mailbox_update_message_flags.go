package imapbackend

import (
	"strings"

	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend/backendutil"
)

func (self *Mailbox) UpdateMessagesFlags(uid bool, seqSet *imap.SeqSet, op imap.FlagsOp, flags []string) error {
	logger.Tracef("Mailbox:UpdateMessagesFlags()")

	return self.backend.db.Update(func(tx *genji.Tx) error {

		var seq uint32 = 0
		return self.backend.mailboxMessagesDAO.FindTx(tx, nil, 0, func(mailboxMessage *model.MailboxMessage) error {

			seq++

			// skip messages that don't match seqSet
			if uid {
				if !seqSet.Contains(mailboxMessage.Uid) {
					return nil
				}
			} else {
				if !seqSet.Contains(seq) {
					return nil
				}
			}

			existingFlags := strings.Split(mailboxMessage.FlagsCSV, ",")
			newFlags := backendutil.UpdateFlags(existingFlags, op, flags)
			mailboxMessage.FlagsCSV = strings.Join(newFlags, ",")
			return self.backend.mailboxMessagesDAO.ReplaceTx(tx, mailboxMessage)
		})
	})
}
