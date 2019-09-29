package imapbackend

import (
	"bytes"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
)

func (self *Mailbox) ListMessages(uid bool, seqSet *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	logger.Tracef("Mailbox:ListMessages()")

	return self.backend.db.View(func(tx *genji.Tx) error {

		var seq uint32 = 0
		return self.backend.mailboxMessagesDAO.FindTx(tx, nil, 0, func(mailboxMessage *model.MailboxMessage) error {

			seq++

			// filter messages that don't match seqSet
			if uid {
				if !seqSet.Contains(mailboxMessage.Uid) {
					return nil
				}
			} else {
				if !seqSet.Contains(seq) {
					return nil
				}
			}

			imapMessage := imap.NewMessage(seq, items)

			// populate requested items
			var message *model.Message // load on-demand
			var messageRawBody *model.MessageRawBody
			for _, item := range items {
				switch item {

				case imap.FetchRFC822, "BODY[]":
					if messageRawBody == nil {
						var err error
						messageRawBody, err = self.backend.messageRawBodiesDAO.FindById(mailboxMessage.MessageId)
						if err != nil {
							return err
						}
					}
					if messageRawBody != nil {
						imapMessage.Items[item] = messageRawBody.Body
					} else {
						logger.Warnf("Body expected but not found: message id %+v", mailboxMessage.MessageId)
						imapMessage.Items[item] = []byte{}
					}

				case imap.FetchEnvelope:
					if messageRawBody == nil {
						messageRawBody, err := self.backend.messageRawBodiesDAO.FindById(mailboxMessage.MessageId)
						if err != nil {
							return err
						}
						envelope := &imap.Envelope{}
						parsedMessage, err := mail.ReadMessage(bytes.NewReader(messageRawBody.Body))
						if err != nil {
							logger.Warnf("Failed parsing message %s: %v", mailboxMessage.MessageId, err)
						} else {
							envelope.Subject = parsedMessage.Header.Get("Subject")
							if date, err := parsedMessage.Header.Date(); err == nil {
								envelope.Date = date
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "From"); err == nil {
								envelope.From = imapAddresses
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "Sender"); err == nil {
								envelope.Sender = imapAddresses
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "Reply-To"); err == nil {
								envelope.ReplyTo = imapAddresses
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "To"); err == nil {
								envelope.To = imapAddresses
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "Cc"); err == nil {
								envelope.Cc = imapAddresses
							}
							if imapAddresses, err := toImapAddressList(parsedMessage, "Bcc"); err == nil {
								envelope.Bcc = imapAddresses
							}
							envelope.InReplyTo = parsedMessage.Header.Get("In-Reply-To")
							envelope.MessageId = parsedMessage.Header.Get("Message-ID")
						}
						imapMessage.Envelope = envelope
					}

				case imap.FetchFlags:
					imapMessage.Items[item] = strings.Split(mailboxMessage.FlagsCSV, ",")

				case imap.FetchInternalDate:
					if message == nil {
						var err error
						message, err = self.backend.messagesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
						if err != nil {
							return err
						}
					}
					imapMessage.InternalDate = time.Unix(message.DateUTC, 0)

				case imap.FetchRFC822Size:
					if message == nil {
						var err error
						message, err = self.backend.messagesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
						if err != nil {
							return err
						}
					}
					imapMessage.Size = message.Size

				case imap.FetchUid:
					imapMessage.Uid = mailboxMessage.Uid

				default:
					return errors.New(fmt.Sprintf("Not implemented yet: unsupported fetch item %s", item))
				}
			}

			ch <- imapMessage
			return nil
		})
	})
}

func toImapAddressList(parsedMessage *mail.Message, headerName string) ([]*imap.Address, error) {
	addresses, err := parsedMessage.Header.AddressList(headerName)
	if err != nil {
		return nil, err
	}
	imapAddressList := []*imap.Address{}
	for _, address := range addresses {
		imapAddress := &imap.Address{PersonalName: address.Name}
		if atIndex := strings.Index(address.Address, "@"); atIndex == -1 {
			imapAddress.MailboxName = address.Address
		} else {
			imapAddress.MailboxName = address.Address[:atIndex]
			imapAddress.HostName = address.Address[atIndex+1:]
		}
		imapAddressList = append(imapAddressList, imapAddress)
	}
	return imapAddressList, nil
}
