package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
)

func (self *User) RenameMailbox(existingName, newName string) error {
	logger.Tracef("User:RenameMailbox()")
	return errors.New("NIY")
}
