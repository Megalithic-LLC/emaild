package model

import (
	"github.com/asdine/genji/record"
)

const (
	AccountTable         = "a"
	DomainTable          = "d"
	MailboxTable         = "mbx"
	MailboxMessageTable  = "mbx_msg"
	MessageRawBodyTable  = "msg_braw"
	MessageTable         = "msg"
	PropertyTable        = "prop"
	ServiceInstanceTable = "si"
	SnapshotTable        = "s"
)

var (
	Tables = map[string]record.Record{
		AccountTable:         new(Account),
		DomainTable:          new(Domain),
		MailboxTable:         new(Mailbox),
		MailboxMessageTable:  new(MailboxMessage),
		MessageRawBodyTable:  new(MessageRawBody),
		MessageTable:         new(Message),
		PropertyTable:        new(Property),
		ServiceInstanceTable: new(ServiceInstance),
		SnapshotTable:        new(Snapshot),
	}
)
