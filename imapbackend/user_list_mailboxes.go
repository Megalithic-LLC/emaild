package imapbackend

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/emersion/go-imap/backend"
)

func (self *User) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	logger.Tracef("User:ListMailboxes()")
	mailboxBackends := []backend.Mailbox{}
	err := self.backend.db.View(func(tx *genji.Tx) error {
		table, err := tx.GetTable("mailboxes")
		if err != nil {
			return err
		}
		//fields := model.NewMailboxFields()
		return query.
			Select().
			From(table).
			//Where(fields.Account.Eq(user.ID)).
			Run(tx).
			Iterate(func(_ []byte, r record.Record) error {
				var mailbox model.Mailbox
				if err := mailbox.ScanRecord(r); err != nil {
					return err
				}
				mailboxBackend := Mailbox{
					backend: self.backend,
					model:   &mailbox,
					name:    mailbox.Name,
					user:    self,
				}

				mailboxBackends = append(mailboxBackends, &mailboxBackend)
				return nil
			})
	})
	logger.Debugf("* returning: %+v", mailboxBackends)
	return mailboxBackends, err
}
