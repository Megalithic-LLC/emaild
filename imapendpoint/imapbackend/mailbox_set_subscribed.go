package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *Mailbox) SetSubscribed(subscribed bool) error {
	logger.Tracef("Mailbox:SetSubscribed(%v)", subscribed)
	self.model.Subscribed = subscribed
	return self.backend.mailboxesDAO.Replace(self.model)
}
