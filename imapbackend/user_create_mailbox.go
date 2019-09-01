package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
)

func (self *User) CreateMailbox(name string) error {
	logger.Tracef("User:CreateMailbox()")
	return errors.New("NIY")
}
