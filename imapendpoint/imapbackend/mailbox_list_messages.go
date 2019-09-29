package imapbackend

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) ListMessages(uid bool, seqSet *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	logger.Tracef("Mailbox:ListMessages()")

	return self.backend.db.View(func(tx *genji.Tx) error {

		var seq uint32 = 0
		return self.backend.mailboxMessagesDAO.FindTx(tx, nil, 0, func(mailboxMessage *model.MailboxMessage) error {

			seq++

			// filter messages that don't match seqSet
			if uid {
				if !seqSet.Contains(mailboxMessage.UID) {
					return nil
				}
			} else {
				if !seqSet.Contains(seq) {
					return nil
				}
			}

			imapMessage := imap.NewMessage(seq, items)

			// populate requested items
			for _, item := range items {
				switch item {
				case imap.FetchRFC822, "BODY[]":
					messageRawBody, err := self.backend.messageRawBodiesDAO.FindByID(mailboxMessage.MessageID)
					if err != nil {
						return err
					}
					if messageRawBody != nil {
						imapMessage.Items[item] = messageRawBody.Body
					} else {
						logger.Warnf("Body expected but not found: message id %+v", mailboxMessage.MessageID)
						imapMessage.Items[item] = []byte{}
					}
				case imap.FetchFlags:
					imapMessage.Items[item] = strings.Split(mailboxMessage.FlagsCSV, ",")
				case imap.FetchInternalDate:
					message, err := self.backend.messagesDAO.FindByIDTx(tx, mailboxMessage.MessageID)
					if err != nil {
						return err
					}
					imapMessage.InternalDate = time.Unix(message.DateUTC, 0)
				case imap.FetchUid:
					imapMessage.Uid = mailboxMessage.UID
				default:
					return errors.New(fmt.Sprintf("Not implemented yet: unsupported fetch item %s", item))
				}
			}

			ch <- imapMessage
			return nil
		})
	})
}
