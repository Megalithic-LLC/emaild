package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/docktermj/go-logger/logger"
)

func (self *User) CreateMailbox(name string) error {
	logger.Tracef("User:CreateMailbox(%s)", name)
	mailbox := &model.Mailbox{
		AccountId:   self.account.Id,
		Name:        name,
		UidValidity: 1,
	}
	return self.backend.mailboxesDAO.Create(mailbox)
}
