package model

import (
	"github.com/asdine/genji/field"
)

type MailboxMessage struct {
	MailboxId string `genji:"index"`
	MessageId string `genji:"index"`
	Uid       uint32 `genji:"index"`
	FlagsCSV  string
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (self *MailboxMessage) PrimaryKey() ([]byte, error) {
	return field.EncodeString(self.MailboxId + ":" + self.MessageId), nil
}
