package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/asdine/genji"
)

type ImapBackend struct {
	accountsDAO dao.AccountsDAO
	db          *genji.DB
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
) *ImapBackend {
	self := ImapBackend{
		accountsDAO: accountsDAO,
		db:          db,
	}
	return &self
}
