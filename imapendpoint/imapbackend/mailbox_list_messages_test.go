package imapbackend_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/emersion/go-imap"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMailboxListMessages(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	var genjiEngine *engine.Engine
	var db *genji.DB
	var imapBackend *imapbackend.ImapBackend
	var accountsDAO dao.AccountsDAO
	var mailboxesDAO dao.MailboxesDAO
	var mailboxMessagesDAO dao.MailboxMessagesDAO
	var messageRawBodiesDAO dao.MessageRawBodiesDAO
	var messagesDAO dao.MessagesDAO

	g.Describe("Mailbox", func() {

		g.BeforeEach(func() {
			genjiEngine = newGenjiEngine()
			db = newDB(genjiEngine)
			accountsDAO = dao.NewAccountsDAO(db)
			mailboxesDAO = dao.NewMailboxesDAO(db)
			mailboxMessagesDAO = dao.NewMailboxMessagesDAO(db)
			messageRawBodiesDAO = dao.NewMessageRawBodiesDAO(db)
			messagesDAO = dao.NewMessagesDAO(db)
			imapBackend = newImapBackend(accountsDAO, db, mailboxesDAO, mailboxMessagesDAO, messageRawBodiesDAO, messagesDAO)
		})
		g.AfterEach(func() {
			closeAndDestroyGenjiEngine(genjiEngine)
		})

		g.Describe("ListMessages()", func() {

			g.It("Should correctly return internal dates", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())
				time1 := time.Date(2010, 2, 20, 22, 7, 15, 0, time.UTC)
				time2 := time.Date(2011, 3, 21, 23, 8, 16, 0, time.UTC)
				time3 := time.Date(2012, 4, 22, 24, 9, 17, 0, time.UTC)
				Expect(mailbox.CreateMessage([]string{}, time1, strings.NewReader("Subject: A\r\n\r\nBody_A"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time2, strings.NewReader("Subject: B\r\n\r\nBody_B"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time3, strings.NewReader("Subject: C\r\n\r\nBody_C"))).ToNot(HaveOccurred())

				// Perform test

				messages := make(chan *imap.Message, 3)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 3)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchInternalDate)}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message1, message2, message3 := <-messages, <-messages, <-messages

				Expect(message1.InternalDate.UTC()).To(Equal(time1.UTC()))
				Expect(message2.InternalDate.UTC()).To(Equal(time2.UTC()))
				Expect(message3.InternalDate.UTC()).To(Equal(time3.UTC()))
			})

			g.It("Should correctly return envelopes", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())

				// message #1 tests maximal absense of fields
				raw1 := "foo: bar\r\n\r\nBody_A"

				// message #2 tests fields populated with a single instance
				raw2 := "Subject: B\r\n" +
					"Date: Wed, 02 Oct 2002 22:08:16 GMT\r\n" +
					"From: \"Joe\" <joe@acme.org>\r\n" +
					"Sender: \"Mary\" <mary@acme.org>\r\n" +
					"Reply-To: \"Fred\" <fred@acme.org>\r\n" +
					"To: \"Serge\" <serge@acme.org>\r\n" +
					"Cc: \"Lisa\" <lisa@acme.org>\r\n" +
					"Bcc: \"Sue\" <sue@acme.org>\r\n" +
					"In-Reply-To: abc123\r\n" +
					"Message-ID: def345\r\n" +
					"\r\n" +
					"Body_B"

				// message #3 tests fields populated with multiple instances where possible
				raw3 := "Subject: C\r\n" +
					"Date: Thu, 03 Oct 2002 23:09:17 GMT\r\n" +
					"From: \"Joe\" <joe@acme.org>, \"Joe2\" <joe2@acme.org>\r\n" +
					"Sender: \"Mary\" <mary@acme.org>, \"Mary2\" <mary2@acme.org>\r\n" +
					"Reply-To: \"Fred\" <fred@acme.org>, \"Fred2\" <fred2@acme.org>\r\n" +
					"To: \"Serge\" <serge@acme.org>, \"Serge2\" <serge2@acme.org>\r\n" +
					"Cc: \"Lisa\" <lisa@acme.org>, \"Lisa2\" <lisa2@acme.org>\r\n" +
					"Bcc: \"Sue\" <sue@acme.org>, \"Sue2\" <sue2@acme.org>\r\n" +
					"In-Reply-To: abc123\r\n" +
					"Message-ID: def345\r\n" +
					"\r\n" +
					"Body_C"

				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw1))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw2))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw3))).ToNot(HaveOccurred())

				// Perform test

				messages := make(chan *imap.Message, 3)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 3)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchEnvelope)}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message1, message2, message3 := <-messages, <-messages, <-messages

				Expect(message1.Envelope).ToNot(BeNil())

				Expect(message2.Envelope).ToNot(BeNil())
				Expect(message2.Envelope.Subject).To(Equal("B"))
				Expect(message2.Envelope.Date.UTC()).To(Equal(time.Date(2002, 10, 2, 22, 8, 16, 0, time.UTC).UTC()))
				Expect(message2.Envelope.From).ToNot(BeNil())
				Expect(message2.Envelope.From).To(HaveLen(1))
				Expect(message2.Envelope.From).To(ContainElement(&imap.Address{PersonalName: "Joe", MailboxName: "joe", HostName: "acme.org"}))
				Expect(message2.Envelope.Sender).To(HaveLen(1))
				Expect(message2.Envelope.Sender).To(ContainElement(&imap.Address{PersonalName: "Mary", MailboxName: "mary", HostName: "acme.org"}))
				Expect(message2.Envelope.ReplyTo).To(HaveLen(1))
				Expect(message2.Envelope.ReplyTo).To(ContainElement(&imap.Address{PersonalName: "Fred", MailboxName: "fred", HostName: "acme.org"}))
				Expect(message2.Envelope.To).To(HaveLen(1))
				Expect(message2.Envelope.To).To(ContainElement(&imap.Address{PersonalName: "Serge", MailboxName: "serge", HostName: "acme.org"}))
				Expect(message2.Envelope.Cc).To(HaveLen(1))
				Expect(message2.Envelope.Cc).To(ContainElement(&imap.Address{PersonalName: "Lisa", MailboxName: "lisa", HostName: "acme.org"}))
				Expect(message2.Envelope.Bcc).To(HaveLen(1))
				Expect(message2.Envelope.Bcc).To(ContainElement(&imap.Address{PersonalName: "Sue", MailboxName: "sue", HostName: "acme.org"}))
				Expect(message2.Envelope.InReplyTo).To(Equal("abc123"))
				Expect(message2.Envelope.MessageId).To(Equal("def345"))

				Expect(message3.Envelope).ToNot(BeNil())
				Expect(message3.Envelope.Subject).To(Equal("C"))
				Expect(message3.Envelope.Date.UTC()).To(Equal(time.Date(2002, 10, 3, 23, 9, 17, 0, time.UTC).UTC()))
				Expect(message3.Envelope.From).ToNot(BeNil())
				Expect(message3.Envelope.From).To(HaveLen(2))
				Expect(message3.Envelope.From).To(ContainElement(&imap.Address{PersonalName: "Joe", MailboxName: "joe", HostName: "acme.org"}))
				Expect(message3.Envelope.From).To(ContainElement(&imap.Address{PersonalName: "Joe2", MailboxName: "joe2", HostName: "acme.org"}))
				Expect(message3.Envelope.Sender).To(HaveLen(2))
				Expect(message3.Envelope.Sender).To(ContainElement(&imap.Address{PersonalName: "Mary", MailboxName: "mary", HostName: "acme.org"}))
				Expect(message3.Envelope.Sender).To(ContainElement(&imap.Address{PersonalName: "Mary2", MailboxName: "mary2", HostName: "acme.org"}))
				Expect(message3.Envelope.ReplyTo).To(HaveLen(2))
				Expect(message3.Envelope.ReplyTo).To(ContainElement(&imap.Address{PersonalName: "Fred", MailboxName: "fred", HostName: "acme.org"}))
				Expect(message3.Envelope.ReplyTo).To(ContainElement(&imap.Address{PersonalName: "Fred2", MailboxName: "fred2", HostName: "acme.org"}))
				Expect(message3.Envelope.To).To(HaveLen(2))
				Expect(message3.Envelope.To).To(ContainElement(&imap.Address{PersonalName: "Serge", MailboxName: "serge", HostName: "acme.org"}))
				Expect(message3.Envelope.To).To(ContainElement(&imap.Address{PersonalName: "Serge2", MailboxName: "serge2", HostName: "acme.org"}))
				Expect(message3.Envelope.Cc).To(HaveLen(2))
				Expect(message3.Envelope.Cc).To(ContainElement(&imap.Address{PersonalName: "Lisa", MailboxName: "lisa", HostName: "acme.org"}))
				Expect(message3.Envelope.Cc).To(ContainElement(&imap.Address{PersonalName: "Lisa2", MailboxName: "lisa2", HostName: "acme.org"}))
				Expect(message3.Envelope.Bcc).To(HaveLen(2))
				Expect(message3.Envelope.Bcc).To(ContainElement(&imap.Address{PersonalName: "Sue", MailboxName: "sue", HostName: "acme.org"}))
				Expect(message3.Envelope.Bcc).To(ContainElement(&imap.Address{PersonalName: "Sue2", MailboxName: "sue2", HostName: "acme.org"}))
				Expect(message3.Envelope.InReplyTo).To(Equal("abc123"))
				Expect(message3.Envelope.MessageId).To(Equal("def345"))
			})

			g.It("Should correctly return UIDs", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: A\r\n\r\nBody_A"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: B\r\n\r\nBody_B"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: C\r\n\r\nBody_C"))).ToNot(HaveOccurred())

				// Perform test

				messages := make(chan *imap.Message, 3)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 3)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchUid)}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message1, message2, message3 := <-messages, <-messages, <-messages

				Expect(message1.SeqNum).To(Equal(uint32(1)))
				Expect(message2.SeqNum).To(Equal(uint32(2)))
				Expect(message3.SeqNum).To(Equal(uint32(3)))

				Expect(message2.Uid).To(Equal(message1.Uid + 1))
				Expect(message3.Uid).To(Equal(message2.Uid + 1))
			})

			g.It("Should correctly return sizes", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: A\r\n\r\n1"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: B\r\n\r\n22"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: C\r\n\r\n333"))).ToNot(HaveOccurred())

				// Perform test

				messages := make(chan *imap.Message, 3)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 3)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchRFC822Size)}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message1, message2, message3 := <-messages, <-messages, <-messages

				Expect(message1.Size).To(Equal(uint32(15)))
				Expect(message2.Size).To(Equal(uint32(16)))
				Expect(message3.Size).To(Equal(uint32(17)))
			})

			g.It("Should return correct bodystructure for a text/plain message", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader("Subject: A\r\n\r\nBodyA"))).ToNot(HaveOccurred())

				// Perform test
				messages := make(chan *imap.Message, 1)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 1)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchBodyStructure)}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message := <-messages
				Expect(message.BodyStructure).ToNot(BeNil())
				Expect(message.BodyStructure.MIMEType).To(Equal("text"))
				Expect(message.BodyStructure.MIMESubType).To(Equal("plain"))
			})

			g.It("Should return requested headers", func() {
				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())

				// Perform test positive test
				{
					flags := []string{}
					var date time.Time
					body := "A: a\r\nB: b\r\nC: c\r\n\r\nbody"
					Expect(mailbox.CreateMessage(flags, date, strings.NewReader(body))).ToNot(HaveOccurred())
				}
				messages := make(chan *imap.Message, 1)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 1)
				fetchItem := imap.FetchItem("BODY[HEADER.FIELDS (A C)]")
				items := []imap.FetchItem{fetchItem}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message := <-messages
				Expect(message.SeqNum).To(Equal(uint32(1)))
				Expect(message.Body).ToNot(BeNil())
				section, err := imap.ParseBodySectionName(fetchItem)
				Expect(err).ToNot(HaveOccurred())
				Expect(message.Body).To(HaveKeyWithValue(section, bytes.NewBufferString("A: a\r\nC: c\r\n\r\n")))

				// Perform negative test
				fetchItem = imap.FetchItem("BODY[HEADER.FIELDS.NOT (A C)]")
				items = []imap.FetchItem{fetchItem}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message = <-messages
				Expect(message.SeqNum).To(Equal(uint32(1)))
				Expect(message.Body).ToNot(BeNil())
				section, err = imap.ParseBodySectionName(fetchItem)
				Expect(err).ToNot(HaveOccurred())
				Expect(message.Body).To(HaveKeyWithValue(section, bytes.NewBufferString("B: b\r\n\r\n")))
			})

		})
	})
}
