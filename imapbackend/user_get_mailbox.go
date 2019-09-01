package imapbackend

import (
	"errors"

	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/emersion/go-imap/backend"
)

func (self *User) GetMailbox(name string) (backend.Mailbox, error) {
	logger.Tracef("User:GetMailbox(%s)", name)

	var mailbox *model.Mailbox
	err := self.backend.db.View(func(tx *genji.Tx) error {
		table, err := tx.GetTable(model.MailboxTable)
		if err != nil {
			return err
		}
		fields := model.NewMailboxFields()
		return query.
			Select().
			From(table).
			Where(fields.Name.Eq(name)).
			Limit(1).
			Run(tx).
			Iterate(func(_ []byte, r record.Record) error {
				var m model.Mailbox
				if err := m.ScanRecord(r); err != nil {
					return err
				}
				mailbox = &m
				return nil
			})
	})
	if err != nil {
		return nil, err
	}

	// Create an INBOX when one does not yet exist
	if mailbox == nil && name == "INBOX" {
		if err := self.CreateMailbox(name); err != nil {
			logger.Errorf("Failed creating inbox: %v", err)
			return nil, err
		}
	}

	if mailbox == nil {
		return nil, errors.New("No such mailbox")
	}

	// Return a mailbox backend adapter
	mailboxBackend := Mailbox{
		backend: self.backend,
		model:   mailbox,
		name:    name,
		user:    self,
	}
	return &mailboxBackend, nil
}
