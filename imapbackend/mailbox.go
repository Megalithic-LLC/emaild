package imapbackend

import (
	"github.com/emersion/go-imap"
)

const (
	delimiter = "/"
)

type Mailbox struct {
	backend *ImapBackend
	name    string
	user    *User
}

func (self *Mailbox) Info() (*imap.MailboxInfo, error) {
	info := &imap.MailboxInfo{
		Delimiter: delimiter,
		Name:      self.name,
	}
	return info, nil
}

func (self *Mailbox) Name() string {
	return self.name
}
