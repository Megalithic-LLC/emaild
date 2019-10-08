package submissionbackend

import (
	"github.com/on-prem-net/emaild/model"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-smtp"
)

func (self *SubmissionBackend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	logger.Tracef("SubmissionBackend:Login(%s)", username)

	// Verify account
	account, err := self.accountsDAO.FindOneByUsername(username)
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
