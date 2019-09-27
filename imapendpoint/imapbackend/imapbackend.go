package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/asdine/genji"
)

type ImapBackend struct {
	accountsDAO  dao.AccountsDAO
	db           *genji.DB
	mailboxesDAO dao.MailboxesDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	mailboxesDAO dao.MailboxesDAO,
) *ImapBackend {
	self := ImapBackend{
		accountsDAO:  accountsDAO,
		db:           db,
		mailboxesDAO: mailboxesDAO,
	}
	return &self
}
