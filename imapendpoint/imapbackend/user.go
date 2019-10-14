package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
)

type User struct {
	account *model.Account
	backend *ImapBackend
}

func (self *User) Username() string {
	return self.account.Email
}
