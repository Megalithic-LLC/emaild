package smtpbackend

import (
	"errors"
	"net/mail"

	"github.com/docktermj/go-logger/logger"
)

func (self *Session) Rcpt(to string) error {
	logger.Tracef("Session:Rcpt(%s)", to)

	if _, err := mail.ParseAddress(to); err != nil {
		return err
	}

	account, err := self.backend.accountsDAO.FindOneByEmail(to)
	if err != nil {
		logger.Errorf("Failure looking up recipient account: %v", err)
		return errors.New("An internal error has occurred")
	}
	if account != nil {
		mailbox, err := self.backend.mailboxesDAO.FindOneByName(account.Id, "INBOX")
		if err != nil {
			logger.Errorf("Failure looking up recipient account: %v", err)
			return errors.New("An internal error has occurred")
		}
		self.recipientMailboxes = append(self.recipientMailboxes, mailbox)
	}

	return nil
}
