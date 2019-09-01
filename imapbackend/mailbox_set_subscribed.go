package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
)

func (self *Mailbox) SetSubscribed(subscribed bool) error {
	logger.Tracef("Mailbox:SetSubscribed(%v)", subscribed)
	return errors.New("NIY")
}
