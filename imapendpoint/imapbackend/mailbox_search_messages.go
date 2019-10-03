package imapbackend

import (
	"bytes"
	"strings"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend/backendutil"
	"github.com/emersion/go-message"
)

func (self *Mailbox) SearchMessages(uid bool, criteria *imap.SearchCriteria) ([]uint32, error) {
	logger.Tracef("Mailbox:SearchMessages()")

	matches := []uint32{}
	err := self.backend.db.View(func(tx *genji.Tx) error {

		var seq uint32 = 0
		return self.backend.mailboxMessagesDAO.FindTx(tx, nil, 0, func(mailboxMessage *model.MailboxMessage) error {
			seq++

			msg, err := self.backend.messagesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
			if err != nil {
				return err
			}

			messageRawBody, err := self.backend.messageRawBodiesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
			if err != nil {
				return err
			}

			parsedMessage, err := message.Read(bytes.NewReader(messageRawBody.Body))
			if err != nil {
				return err
			}

			flags := strings.Split(mailboxMessage.FlagsCSV, ",")
			ok, err := backendutil.Match(parsedMessage, seq, mailboxMessage.Uid, time.Unix(msg.DateUTC, 0), flags, criteria)
			if err != nil {
				return err
			}

			if ok {
				if uid {
					matches = append(matches, mailboxMessage.Uid)
				} else {
					matches = append(matches, seq)
				}
			}
			return nil
		})
	})
	return matches, err
}
