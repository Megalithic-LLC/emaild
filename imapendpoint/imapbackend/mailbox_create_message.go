package imapbackend

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	logger.Tracef("Mailbox:CreateMessage()")
	return self.backend.db.Update(func(tx *genji.Tx) error {
		message := &model.Message{
			DateUTC:  date.UTC().Unix(),
			FlagsCSV: strings.Join(flags, ","),
		}
		if err := self.backend.messagesDAO.CreateTx(tx, message); err != nil {
			logger.Errorf("Failed storing message: %v", err)
			return err
		}

		uid, err := self.backend.mailboxesDAO.AllocateNextUidTx(tx, self.model)
		if err != nil {
			return err
		}
		mailboxMessage := &model.MailboxMessage{
			MailboxID: self.model.ID,
			MessageID: message.ID,
			UID:       uid,
		}
		if err := self.backend.mailboxMessagesDAO.CreateTx(tx, mailboxMessage); err != nil {
			logger.Errorf("Failed storing message in mailbox: %v", err)
			return err
		}

		rawBody, err := ioutil.ReadAll(body)
		if err != nil {
			logger.Errorf("Failed reading uploaded message: %v", err)
			return err
		}
		messageRawBody := &model.MessageRawBody{
			ID:   message.ID,
			Body: rawBody,
		}
		if err := self.backend.messageRawBodiesDAO.CreateTx(tx, messageRawBody); err != nil {
			logger.Errorf("Failed storing message body in mailbox: %v", err)
			return err
		}

		return nil
	})
}
