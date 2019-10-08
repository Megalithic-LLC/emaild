package smtpbackend

import (
	"github.com/on-prem-net/emaild/model"
)

type Session struct {
	account    *model.Account
	backend    *SmtpBackend
	recipients []*model.Account
}
