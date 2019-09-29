package imapbackend

import (
	"strings"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/query"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) Expunge() error {
	logger.Tracef("Mailbox:Expunge()")
	return self.backend.db.Update(func(tx *genji.Tx) error {
		var where query.Expr
		return self.backend.mailboxMessagesDAO.FindTx(tx, where, 0, func(mailboxMessage *model.MailboxMessage) error {
			flags := strings.Split(mailboxMessage.FlagsCSV, ",")
			if flagsContains(flags, imap.DeletedFlag) {
				return self.backend.mailboxMessagesDAO.DeleteTx(tx, mailboxMessage.MailboxId, mailboxMessage.MessageId)
			}
			return nil
		})
	})
}
