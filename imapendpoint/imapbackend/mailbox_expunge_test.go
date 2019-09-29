package imapbackend_test

import (
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

func TestMailboxExpunge(t *testing.T) {
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

	g.Describe("Mailbox operations", func() {

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

		g.Describe("Expunge()", func() {

			g.It("Should work", func() {

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
				Expect(mailbox.CreateMessage([]string{imap.RecentFlag}, time1, strings.NewReader("Subject: A\r\n\r\nBody_A"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{imap.RecentFlag}, time2, strings.NewReader("Subject: B\r\n\r\nBody_B"))).ToNot(HaveOccurred())
				Expect(mailbox.CreateMessage([]string{imap.RecentFlag}, time3, strings.NewReader("Subject: C\r\n\r\nBody_C"))).ToNot(HaveOccurred())

				// TODO Sanity check
				//status, err := mailbox.Status([]imap.StatusItem{imap.StatusMessages, imap.StatusRecent, imap.StatusUnseen})
				//Expect(err).ToNot(HaveOccurred())
				//Expect(status.Messages).To(Equal(3))
				//Expect(status.Recent).To(Equal(3))
				//Expect(status.Unseen).To(Equal(3))

				// Perform test
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(2, 2)
				Expect(mailbox.UpdateMessagesFlags(false, seqSet, imap.AddFlags, []string{imap.DeletedFlag})).ToNot(HaveOccurred())
				Expect(mailbox.Expunge()).ToNot(HaveOccurred())

				// Verify fetch does not return expunged message
				messages := make(chan *imap.Message, 2)
				seqSet = new(imap.SeqSet)
				seqSet.AddRange(1, 2)
				items := []imap.FetchItem{imap.FetchItem(imap.FetchInternalDate)}
				Expect(mailbox.ListMessages(false, seqSet, items, messages)).ToNot(HaveOccurred())
				message1, message2 := <-messages, <-messages
				Expect(message1.InternalDate.UTC()).To(Equal(time1.UTC()))
				Expect(message2.InternalDate.UTC()).To(Equal(time3.UTC()))

				// TODO Verify status does not return expunged message
				//status, err := mailbox.Status([]imap.StatusItem{imap.StatusMessages, imap.StatusRecent, imap.StatusUnseen})
				//Expect(err).ToNot(HaveOccurred())
				//Expect(status.Messages).To(Equal(2))
				//Expect(status.Recent).To(Equal(2))
				//Expect(status.Unseen).To(Equal(2))
			})

		})
	})
}
