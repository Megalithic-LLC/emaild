package imapbackend

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/table"
	"github.com/docktermj/go-logger/logger"
)

func (self *User) DeleteMailbox(name string) error {
	logger.Tracef("User:DeleteMailbox()")

	return self.backend.db.Update(func(tx *genji.Tx) error {
		// Delete mailbox
		mailbox, err := self.backend.mailboxesDAO.FindOneByNameTx(tx, self.account.Id, name)
		if err != nil {
			return err
		}
		if err := self.backend.mailboxesDAO.DeleteTx(tx, mailbox.Id); err != nil {
			return err
		}

		// Delete cross-post (message-to-mailbox association)
		mailboxMessageFields := model.NewMailboxMessageFields()
		where := mailboxMessageFields.MailboxId.Eq(mailbox.Id)
		return self.backend.mailboxMessagesDAO.FindTx(tx, where, 0, func(mailboxMessage *model.MailboxMessage) error {
			if err := self.backend.mailboxMessagesDAO.DeleteTx(tx, mailboxMessage.MailboxId, mailboxMessage.MessageId); err != nil {
				return err
			}

			// Delete a message if it's no longer cross-posted into any other mailboxes
			shouldDeleteMessage := true
			where := mailboxMessageFields.MessageId.Eq(mailboxMessage.MessageId)
			if err := self.backend.mailboxMessagesDAO.FindTx(tx, where, 1, func(_ *model.MailboxMessage) error {
				shouldDeleteMessage = false
				return nil
			}); err != nil {
				return err
			}
			if shouldDeleteMessage {
				if err := self.backend.messageRawBodiesDAO.DeleteTx(tx, mailboxMessage.MessageId); err != nil {
					if err != table.ErrRecordNotFound {
						return err
					}
				}
				if err := self.backend.messagesDAO.DeleteTx(tx, mailboxMessage.MessageId); err != nil {
					if err != table.ErrRecordNotFound {
						return err
					}
				}
			}
			return nil
		})
	})
}
