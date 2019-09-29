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

		g.Describe("ListMessages()", func() {

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

				messages := make(chan *imap.Message, 10)
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

		})
	})
}
