package imapbackend

import (
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
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
			Name:        name,
			UidValidity: 1,
		}
		_, err = table.Insert(mailbox)
		return err
	})
}
