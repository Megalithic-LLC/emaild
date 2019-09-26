package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
)

func (self *User) DeleteMailbox(name string) error {
	logger.Tracef("User:DeleteMailbox()")
	return errors.New("NIY")
}
