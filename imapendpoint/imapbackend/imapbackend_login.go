package imapbackend

import (
	"errors"

	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
)

func (self *ImapBackend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	logger.Tracef("ImapBackend:Login(%s)", username)

	// Verify account
	account, err := self.accountsDAO.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("No such account")
	}

	user := User{backend: self, username: username}
	return &user, nil
}
