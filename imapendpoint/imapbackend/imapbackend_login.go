package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
)

func (self *ImapBackend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	logger.Tracef("ImapBackend:Login(%s)", username)

	// Verify account
	account, err := self.accountsDAO.FindOneByEmail(username)
	if err != nil {
		return nil, err
	}

	user := User{
		account: account,
		backend: self,
	}
	return &user, nil
}
