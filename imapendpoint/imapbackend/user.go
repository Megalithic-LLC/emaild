package imapbackend

import (
	"github.com/on-prem-net/emaild/model"
)

type User struct {
	account *model.Account
	backend *ImapBackend
}

func (self *User) Username() string {
	return self.account.Username
}
