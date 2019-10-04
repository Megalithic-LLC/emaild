package smtpbackend

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Data(r io.Reader) error {
	logger.Tracef("Session:Data()")

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return self.backend.db.Update(func(tx *genji.Tx) error {

		// Store message
		message := &model.Message{
			DateUTC: time.Now().UTC().Unix(),
			Size:    uint32(len(data)),
		}
		if err := self.backend.messagesDAO.CreateTx(tx, message); err != nil {
			return err
		}

		// Store message body
		messageRawBody := &model.MessageRawBody{
			Id:   message.Id,
			Body: data,
		}
		if err := self.backend.messageRawBodiesDAO.CreateTx(tx, messageRawBody); err != nil {
			return err
		}

		// Cross-post message into each recipient mailbox
		for _, mailbox := range self.recipientMailboxes {
			mailboxMessage := &model.MailboxMessage{
				MailboxId: mailbox.Id,
				MessageId: message.Id,
			}
			if err := self.backend.mailboxMessagesDAO.CreateTx(tx, mailboxMessage); err != nil {
				return err
			}
		}

		// Success
		logger.Infof("Message delivered to %d recipients", len(self.recipientMailboxes))
		return nil
	})
}
