package imapbackend

import (
	"errors"
	"strings"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/record"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) UpdateMessagesFlags(uid bool, seqSet *imap.SeqSet, op imap.FlagsOp, flags []string) error {
	logger.Tracef("Mailbox:UpdateMessagesFlags()")

	return self.backend.db.Update(func(tx *genji.Tx) error {

		var seq uint32 = 0
		return self.backend.mailboxMessagesDAO.FindTx(tx, nil, 0, func(recordID []byte, r record.Record) error {
			seq++
			var mailboxMessage model.MailboxMessage
			if err := mailboxMessage.ScanRecord(r); err != nil {
				return err
			}

			// skip messages that don't match seqSet
			if uid {
				if !seqSet.Contains(mailboxMessage.UID) {
					return nil
				}
			} else {
				if !seqSet.Contains(seq) {
					return nil
				}
			}

			switch op {

			// Perform additive setting of new flags
			case imap.AddFlags:
				existingFlags := strings.Split(mailboxMessage.FlagsCSV, ",")
				newFlags := existingFlags
				for _, newFlag := range flags {
					alreadyExists := false
					for _, existingFlag := range existingFlags {
						if existingFlag == newFlag {
							alreadyExists = true
							break
						}
					}
					if !alreadyExists {
						newFlags = append(newFlags, newFlag)
					}
				}
				mailboxMessage.FlagsCSV = strings.Join(newFlags, ",")
				if err := self.backend.mailboxMessagesDAO.ReplaceTx(tx, &mailboxMessage); err != nil {
					return err
				}

			default:
				return errors.New("NIY")

			}

			return nil
		})
	})
}
