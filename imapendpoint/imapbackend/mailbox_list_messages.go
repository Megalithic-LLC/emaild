package imapbackend

import (
	"bufio"
	"bytes"
	"net/mail"
	"strings"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/docktermj/go-logger/logger"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend/backendutil"
	"github.com/emersion/go-message/textproto"
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

				case imap.FetchBody, imap.FetchBodyStructure:
					if messageRawBody == nil {
						var err error
						messageRawBody, err = self.backend.messageRawBodiesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
						if err != nil {
							return err
						}
					}
					bodyReader := bufio.NewReader(bytes.NewReader(messageRawBody.Body))
					header, err := textproto.ReadHeader(bodyReader)
					if err != nil {
						logger.Warnf("Failed parsing message header: %v", err)
					} else {
						bodyStructure, err := backendutil.FetchBodyStructure(header, bodyReader, true)
						if err != nil {
							logger.Warnf("Failed parsing bodystructure: %v", err)
						} else {
							imapMessage.BodyStructure = bodyStructure
						}
					}

				case imap.FetchEnvelope:
					if messageRawBody == nil {
						messageRawBody, err := self.backend.messageRawBodiesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
						if err != nil {
							return err
						}
						bodyReader := bufio.NewReader(bytes.NewReader(messageRawBody.Body))
						header, err := textproto.ReadHeader(bodyReader)
						if err != nil {
							logger.Warnf("Failed parsing message header: %v", err)
						} else {
							imapMessage.Envelope, _ = backendutil.FetchEnvelope(header)
						}
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
					if messageRawBody == nil {
						var err error
						messageRawBody, err = self.backend.messageRawBodiesDAO.FindByIdTx(tx, mailboxMessage.MessageId)
						if err != nil {
							return err
						}
					}

					if messageRawBody == nil {
						logger.Warnf("Body expected but not found: message id %+v", mailboxMessage.MessageId)
						imapMessage.Items[item] = []byte{}
					} else {
						section, err := imap.ParseBodySectionName(item)
						if err != nil {
							return err
						}
						bodyReader := bufio.NewReader(bytes.NewReader(messageRawBody.Body))
						header, err := textproto.ReadHeader(bodyReader)
						if err != nil {
							logger.Warnf("Failed parsing message header: %v", err)
						} else {
							data, _ := backendutil.FetchBodySection(header, bodyReader, section)
							imapMessage.Body[section] = data
						}
					}

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
