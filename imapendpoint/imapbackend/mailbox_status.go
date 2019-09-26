package imapbackend

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) Status(items []imap.StatusItem) (*imap.MailboxStatus, error) {
	logger.Tracef("Mailbox:Status()")

	status := imap.NewMailboxStatus(self.name, items)
	//status.Flags = self.flags()
	status.PermanentFlags = []string{"\\*"}
	//status.UnseenSeqNum = self.unseenSeqNum()

	for _, itemName := range items {
		switch itemName {
		case imap.StatusMessages:
			status.Messages = self.model.Messages
		case imap.StatusUidNext:
			status.UidNext = self.model.UidNext
		case imap.StatusUidValidity:
			status.UidValidity = self.model.UidValidity
		case imap.StatusRecent:
			status.Recent = self.model.Recent
		case imap.StatusUnseen:
			status.Unseen = self.model.Unseen
		}
	}

	return status, nil
}
