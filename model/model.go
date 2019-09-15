package model

import (
	"github.com/asdine/genji/record"
)

const (
	AccountTable        = "acct"
	MailboxTable        = "mbx"
	MailboxMessageTable = "mbx_msg"
	MessageTable        = "msg"
	MessageBodyRawTable = "msg_braw"
	PropertyTable       = "prop"
)

var (
	Tables = map[string]record.Record{
		AccountTable:        new(Account),
		MailboxTable:        new(Mailbox),
		MailboxMessageTable: new(MailboxMessage),
		MessageTable:        new(Message),
		MessageBodyRawTable: new(MessageBodyRaw),
		PropertyTable:       new(Property),
	}
)
