package imapbackend

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	logger.Tracef("Mailbox:CreateMessage()")

	return self.backend.db.Update(func(tx *genji.Tx) error {

		// Freshen model
		mailbox, err := self.backend.mailboxesDAO.FindByIdTx(tx, self.model.Id)
		if err != nil {
			return err
		}

		rawBody, err := ioutil.ReadAll(body)
		if err != nil {
			logger.Errorf("Failed reading uploaded message: %v", err)
			return err
		}

		// Store message
		message := &model.Message{
			DateUTC: date.UTC().Unix(),
			Size:    uint32(len(rawBody)),
		}
		if err := self.backend.messagesDAO.CreateTx(tx, message); err != nil {
			logger.Errorf("Failed storing message: %v", err)
			return err
		}

		// Store message body
		messageRawBody := &model.MessageRawBody{
			Id:   message.Id,
			Body: rawBody,
		}
		if err := self.backend.messageRawBodiesDAO.CreateTx(tx, messageRawBody); err != nil {
			logger.Errorf("Failed storing message body in mailbox: %v", err)
			return err
		}

		// Store mailbox-to-message cross-post association, where uid and most flags are stored
		uid, err := self.backend.mailboxesDAO.AllocateNextUidTx(tx, mailbox)
		if err != nil {
			return err
		}
		mailboxMessage := &model.MailboxMessage{
			MailboxId: self.model.Id,
			MessageId: message.Id,
			Uid:       uid,
			FlagsCSV:  strings.Join(flags, ","),
		}
		if err := self.backend.mailboxMessagesDAO.CreateTx(tx, mailboxMessage); err != nil {
			logger.Errorf("Failed storing message in mailbox: %v", err)
			return err
		}

		// Success
		self.model = mailbox
		return nil
	})
}
