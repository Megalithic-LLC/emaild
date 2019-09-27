package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
)

func (self *User) GetMailbox(name string) (backend.Mailbox, error) {
	logger.Tracef("User:GetMailbox(%s)", name)

	mailbox, err := self.backend.mailboxesDAO.FindOneByName(self.account.ID, name)
	if err != nil {
		return nil, err
	}

	// Create an INBOX when one does not yet exist
	if mailbox == nil && name == "INBOX" {
		if err := self.CreateMailbox(name); err != nil {
			logger.Errorf("Failed creating inbox: %v", err)
			return nil, err
		}
	}

	if mailbox == nil {
		return nil, errors.New("No such mailbox")
	}

	// Return a mailbox backend adapter
	mailboxBackend := Mailbox{
		backend: self.backend,
		model:   mailbox,
		name:    name,
		user:    self,
	}
	return &mailboxBackend, nil
}
