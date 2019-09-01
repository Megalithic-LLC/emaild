package imapbackend

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	logger.Tracef("Mailbox:CreateMessage()")
	return self.backend.db.Update(func(tx *genji.Tx) error {

		messageTable, err := tx.GetTable(model.MessageTable)
		if err != nil {
			logger.Errorf("Failure: %v", err)
			return err
		}
		message := &model.Message{
			DateUTC:  date.UTC().Unix(),
			FlagsCSV: strings.Join(flags, ","),
		}
		if _, err = messageTable.Insert(message); err != nil {
			logger.Errorf("Failed storing message: %v", err)
			return err
		}

		mailboxMessageTable, err := tx.GetTable(model.MailboxMessageTable)
		if err != nil {
			logger.Errorf("Failure: %v", err)
			return err
		}
		mailboxMessage := &model.MailboxMessage{
			MailboxID: self.model.ID,
			MessageID: message.ID,
		}
		if _, err = mailboxMessageTable.Insert(mailboxMessage); err != nil {
			logger.Errorf("Failed storing message in mailbox: %v", err)
			return err
		}

		// TODO body
		messageBodyRawTable, err := tx.GetTable(model.MessageBodyRawTable)
		if err != nil {
			logger.Errorf("Failure: %v", err)
			return err
		}
		bodyRaw, err := ioutil.ReadAll(body)
		if err != nil {
			logger.Errorf("Failed reading uploaded message: %v", err)
			return err
		}
		messageBodyRaw := &model.MessageBodyRaw{
			ID:   message.ID,
			Body: bodyRaw,
		}
		if _, err = messageBodyRawTable.Insert(messageBodyRaw); err != nil {
			logger.Errorf("Failed storing message body in mailbox: %v", err)
			return err
		}

		return nil
	})
}
