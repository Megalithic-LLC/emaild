package imapbackend

import (
	"strings"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) Status(items []imap.StatusItem) (*imap.MailboxStatus, error) {
	logger.Tracef("Mailbox:Status()")

	status := imap.NewMailboxStatus(self.name, items)
	status.PermanentFlags = []string{"\\*"}
	//status.Flags = self.flags()
	//status.UnseenSeqNum = self.unseenSeqNum()

	err := self.backend.db.View(func(tx *genji.Tx) error {

		// Freshen the cached mailbox model
		mailbox, err := self.backend.mailboxesDAO.FindById(self.model.Id)
		if err != nil {
			return err
		}

		// Count
		mailboxMessageFields := model.NewMailboxMessageFields()
		where := mailboxMessageFields.MailboxId.Eq(self.model.Id)
		var messageCount, recentCount, unseenCount uint32 = 0, 0, 0
		if err := self.backend.mailboxMessagesDAO.FindTx(tx, where, 0, func(mailboxMessage *model.MailboxMessage) error {
			messageCount++
			flags := strings.Split(mailboxMessage.FlagsCSV, ",")
			if flagsContains(flags, imap.RecentFlag) {
				recentCount++
			}
			if !flagsContains(flags, imap.SeenFlag) {
				unseenCount++
			}
			return nil
		}); err != nil {
			return err
		}

		for _, itemName := range items {
			switch itemName {
			case imap.StatusMessages:
				status.Messages = messageCount
			case imap.StatusUidNext:
				status.UidNext = mailbox.UidNext
			case imap.StatusUidValidity:
				status.UidValidity = mailbox.UidValidity
			case imap.StatusRecent:
				status.Recent = recentCount
			case imap.StatusUnseen:
				status.Unseen = unseenCount
			}
		}

		// Success
		self.model = mailbox
		return nil
	})

	return status, err
}
