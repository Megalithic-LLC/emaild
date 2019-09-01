package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap/backend"
)

func (self *User) GetMailbox(name string) (backend.Mailbox, error) {
	logger.Tracef("User:GetMailbox(%s)", name)
	mailbox := Mailbox{
		backend: self.backend,
		name:    name,
		user:    self,
	}
	return &mailbox, nil
}
