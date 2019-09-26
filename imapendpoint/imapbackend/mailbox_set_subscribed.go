package imapbackend

import (
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
)

func (self *Mailbox) SetSubscribed(subscribed bool) error {
	logger.Tracef("Mailbox:SetSubscribed(%v)", subscribed)
	return self.backend.db.Update(func(tx *genji.Tx) error {
		table, err := tx.GetTable("mailboxes")
		if err != nil {
			return err
		}
		self.model.Subscribed = subscribed
		return table.Replace([]byte(self.model.ID), self.model)
	})
}
