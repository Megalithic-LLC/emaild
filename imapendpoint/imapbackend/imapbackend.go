package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/asdine/genji"
)

type ImapBackend struct {
	accountsDAO         dao.AccountsDAO
	db                  *genji.DB
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	mailboxesDAO dao.MailboxesDAO,
	mailboxMessagesDAO dao.MailboxMessagesDAO,
	messageRawBodiesDAO dao.MessageRawBodiesDAO,
	messagesDAO dao.MessagesDAO,
) *ImapBackend {
	self := ImapBackend{
		accountsDAO:         accountsDAO,
		db:                  db,
		mailboxesDAO:        mailboxesDAO,
		mailboxMessagesDAO:  mailboxMessagesDAO,
		messageRawBodiesDAO: messageRawBodiesDAO,
		messagesDAO:         messagesDAO,
	}
	return &self
}
