package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/docktermj/go-logger/logger"
)

func (self *User) CreateMailbox(name string) error {
	logger.Tracef("User:CreateMailbox(%s)", name)
	mailbox := model.NewMailbox(self.account.Id, name)
	return self.backend.mailboxesDAO.Create(mailbox)
}
