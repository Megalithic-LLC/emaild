package model

import (
	"github.com/asdine/genji/record"
)

const (
	AccountTable         = "a"
	DomainTable          = "d"
	EndpointTable        = "e"
	MailboxTable         = "mbx"
	MailboxMessageTable  = "mm"
	MessageRawBodyTable  = "mrb"
	MessageTable         = "m"
	PropertyTable        = "p"
	ServiceInstanceTable = "si"
	SnapshotTable        = "s"
)

var (
	Tables = map[string]record.Record{
		AccountTable:         new(Account),
		DomainTable:          new(Domain),
		EndpointTable:        new(Endpoint),
		MailboxTable:         new(Mailbox),
		MailboxMessageTable:  new(MailboxMessage),
		MessageRawBodyTable:  new(MessageRawBody),
		MessageTable:         new(Message),
		PropertyTable:        new(Property),
		ServiceInstanceTable: new(ServiceInstance),
		SnapshotTable:        new(Snapshot),
	}
)
