package imapbackend

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/rs/xid"
)

func (self *User) CreateMailbox(name string) error {
	logger.Tracef("User:CreateMailbox(%s)", name)
	return self.backend.db.Update(func(tx *genji.Tx) error {
		table, err := tx.GetTable(model.MailboxTable)
		if err != nil {
			return err
		}
		mailbox := &model.Mailbox{
			ID:          xid.New().String(),
			AccountID:   self.account.ID,
			Name:        name,
			UidValidity: 1,
		}
		_, err = table.Insert(mailbox)
		return err
	})
}
