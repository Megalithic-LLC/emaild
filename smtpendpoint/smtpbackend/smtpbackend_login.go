package smtpbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

func (self *SmtpBackend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	logger.Tracef("SmtpBackend:Login(%s)", username)

	// Verify account
	account, err := self.accountsDAO.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}

	session := Session{
		account:            account,
		backend:            self,
		recipientMailboxes: []*model.Mailbox{},
	}
	return &session, nil
}
