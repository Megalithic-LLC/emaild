package smtpbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type Session struct {
	account    *model.Account
	backend    *SmtpBackend
	recipients []*model.Account
}
