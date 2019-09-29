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

func TestMailboxCreateMessage(t *testing.T) {
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

		g.Describe("CreateMessage()", func() {

			g.It("Should correctly store a simple text/plain message", func() {
				// setup precondition
				account := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test", "password")
				Expect(err).ToNot(HaveOccurred())
				mailbox, err := user.GetMailbox("INBOX")
				Expect(err).ToNot(HaveOccurred())
				Expect(mailbox).ToNot(BeNil())

				// Perform test
				{
					flags := []string{}
					var date time.Time
					body := "Subject: hi\r\n\r\nbody"
					Expect(mailbox.CreateMessage(flags, date, strings.NewReader(body))).ToNot(HaveOccurred())
				}
				messages := make(chan *imap.Message, 1)
				uid := false
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 1)
				items := []imap.FetchItem{imap.FetchItem("BODY[]")}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message := <-messages
				Expect(message.SeqNum).To(Equal(uint32(1)))
				body := message.Items["BODY[]"]
				Expect(body).ToNot(BeNil())
				Expect(string(body.([]byte))).To(Equal("Subject: hi\r\n\r\nbody"))
			})

		})
	})
}
