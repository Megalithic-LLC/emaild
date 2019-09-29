package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
)

func (self *User) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	logger.Tracef("User:ListMailboxes()")

	fields := model.NewMailboxFields()
	where := fields.Subscribed.Eq(subscribed)

	limit := 0
	mailboxBackends := []backend.Mailbox{}

	err := self.backend.mailboxesDAO.Find(where, limit, func(mailbox *model.Mailbox) error {
		mailboxBackend := Mailbox{
			backend: self.backend,
			model:   mailbox,
			name:    mailbox.Name,
			user:    self,
		}
		mailboxBackends = append(mailboxBackends, &mailboxBackend)
		return nil
	})

	return mailboxBackends, err
}
