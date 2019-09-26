package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
)

func (self *User) Logout() error {
	logger.Tracef("User:Logout()")
	return nil
}
