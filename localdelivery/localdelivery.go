package localdelivery

import (
	"errors"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
)

type LocalDelivery struct {
	accountsDAO         dao.AccountsDAO
	db                  *genji.DB
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	mailboxesDAO dao.MailboxesDAO,
	mailboxMessagesDAO dao.MailboxMessagesDAO,
	messageRawBodiesDAO dao.MessageRawBodiesDAO,
	messagesDAO dao.MessagesDAO,
) *LocalDelivery {
	self := LocalDelivery{
		accountsDAO:         accountsDAO,
		db:                  db,
		mailboxesDAO:        mailboxesDAO,
		mailboxMessagesDAO:  mailboxMessagesDAO,
		messageRawBodiesDAO: messageRawBodiesDAO,
		messagesDAO:         messagesDAO,
	}
	return &self
}

func (self *LocalDelivery) Deliver(data []byte, recipients []*model.Account) error {
	logger.Tracef("LocalDelivery:Deliver()")

	return self.db.Update(func(tx *genji.Tx) error {

		// Store message
		message := &model.Message{
			DateUTC: time.Now().UTC().Unix(),
			Size:    uint32(len(data)),
		}
		if err := self.messagesDAO.CreateTx(tx, message); err != nil {
			return err
		}

		// Store message body
		messageRawBody := &model.MessageRawBody{
			Id:   message.Id,
			Body: data,
		}
		if err := self.messageRawBodiesDAO.CreateTx(tx, messageRawBody); err != nil {
			return err
		}

		// Cross-post message into each recipient mailbox
		for _, recipient := range recipients {

			mailbox, err := self.mailboxesDAO.FindOneByNameTx(tx, recipient.Id, "INBOX")
			if err != nil {
				logger.Errorf("Failure looking up inbox for recipient %v: %v", recipient, err)
				return errors.New("An internal error has occurred")
			}

			// Create an inbox for a recipient if one does not yet exist
			if mailbox == nil {
				mailbox = model.NewMailbox(recipient.Id, "INBOX")
				if err := self.mailboxesDAO.CreateTx(tx, mailbox); err != nil {
					logger.Errorf("Failure creating INBOX for account %v: %v", recipient, err)
					return errors.New("An internal error has occurred")
				}
			}

			// Perform the cross-post
			mailboxMessage := &model.MailboxMessage{
				MailboxId: mailbox.Id,
				MessageId: message.Id,
			}
			if err := self.mailboxMessagesDAO.CreateTx(tx, mailboxMessage); err != nil {
				return err
			}
		}

		// Success
		if len(recipients) == 1 {
			logger.Infof("Message delivered to %s", recipients[0].Email)
		} else {
			logger.Infof("Message delivered to %d recipients", len(recipients))
		}
		return nil
	})
}
