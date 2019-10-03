package imapbackend_test

import (
	"net/textproto"
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

func TestMailboxSearchMessages(t *testing.T) {
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

		g.Describe("SearchMessages()", func() {

			g.It("Should match using From: headers", func() {

				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())

				raw1 := "foo: bar\r\n\r\nBody_A"

				raw2 := "Subject: B\r\n" +
					"Date: Wed, 02 Oct 2002 22:08:16 GMT\r\n" +
					"From: \"Joe\" <joe@acme.org>\r\n" +
					"\r\n" +
					"Body_B"

				raw3 := "Subject: C\r\n" +
					"Date: Thu, 03 Oct 2002 23:09:17 GMT\r\n" +
					"From: \"Joe\" <joe@acme.org>, \"Joe2\" <joe2@acme.org>\r\n" +
					"\r\n" +
					"Body_C"

				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw1))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw2))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{}, time.Now(), strings.NewReader(raw3))).ToNot(HaveOccurred())

				// Perform test

				searchCriteria := &imap.SearchCriteria{
					Header: textproto.MIMEHeader{"From": {"joe@acme.org"}},
				}
				matchingSeqNums, err := mailbox.SearchMessages(false, searchCriteria)
				Expect(err).ToNot(HaveOccurred())
				Expect(matchingSeqNums).To(HaveLen(2))
				Expect(matchingSeqNums).To(ContainElement(uint32(2)))
				Expect(matchingSeqNums).To(ContainElement(uint32(3)))
			})

			g.It("Should match using UIDs", func() {

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
				searchCriteria := &imap.SearchCriteria{
					Uid: &imap.SeqSet{Set: []imap.Seq{{1, 1}, {3, 3}}},
				}
				matchingSeqNums, err := mailbox.SearchMessages(false, searchCriteria)
				Expect(err).ToNot(HaveOccurred())
				Expect(matchingSeqNums).To(HaveLen(2))
				Expect(matchingSeqNums).To(ContainElement(uint32(1)))
				Expect(matchingSeqNums).To(ContainElement(uint32(3)))
			})

			g.It("Should match using sizes", func() {

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

				// FYI larger is currently >= -- https://github.com/emersion/go-imap/issues/298
				searchCriteria := &imap.SearchCriteria{
					Larger: 2,
				}
				matchingSeqNums, err := mailbox.SearchMessages(false, searchCriteria)
				Expect(err).ToNot(HaveOccurred())
				Expect(matchingSeqNums).To(HaveLen(2))
				Expect(matchingSeqNums).To(ContainElement(uint32(2)))
				Expect(matchingSeqNums).To(ContainElement(uint32(3)))

				// FYI smaller is currently <=
				searchCriteria = &imap.SearchCriteria{
					Smaller: 2,
				}
				matchingSeqNums, err = mailbox.SearchMessages(false, searchCriteria)
				Expect(err).ToNot(HaveOccurred())
				Expect(matchingSeqNums).To(HaveLen(2))
				Expect(matchingSeqNums).To(ContainElement(uint32(1)))
				Expect(matchingSeqNums).To(ContainElement(uint32(2)))
			})

		})
	})
}
