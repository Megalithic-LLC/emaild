package smtpbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
	"github.com/on-prem-net/emaild/model"
)

func (self *SmtpBackend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	logger.Tracef("SmtpBackend:Login(%s)", username)

	// Verify account
	account, err := self.accountsDAO.FindOneByEmail(username)
	if err != nil {
		return nil, err
	}

	session := Session{
		account:    account,
		backend:    self,
		recipients: []*model.Account{},
	}
	return &session, nil
}
