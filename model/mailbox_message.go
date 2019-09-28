package model

import (
	"github.com/asdine/genji/field"
)

type MailboxMessage struct {
	MailboxID string `genji:"index"`
	MessageID string `genji:"index"`
	UID       uint32 `genji:"index"`
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (self *MailboxMessage) PrimaryKey() ([]byte, error) {
	return field.EncodeString(self.MailboxID + ":" + self.MessageID), nil
}
