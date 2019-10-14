package imapbackend_test

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/emersion/go-imap"
	"github.com/franela/goblin"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
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

		g.Describe("CreateMessage()", func() {

			g.It("Should correctly store a simple text/plain message", func() {
				// setup precondition
				account := &model.Account{Name: "test", Email: "test@acme.org"}
				Expect(accountsDAO.Create(account)).Should(Succeed())
				user, err := imapBackend.Login(nil, "test@acme.org", "password")
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
				fetchItem := imap.FetchItem("BODY[]")
				items := []imap.FetchItem{fetchItem}
				Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
				message := <-messages
				Expect(message.SeqNum).To(Equal(uint32(1)))
				sectionName, err := imap.ParseBodySectionName(fetchItem)
				Expect(err).ToNot(HaveOccurred())
				section := message.GetBody(sectionName)
				Expect(section).ToNot(BeNil())
				sectionValue, err := ioutil.ReadAll(section)
				Expect(err).ToNot(HaveOccurred())
				Expect(sectionValue).To(Equal([]byte("Subject: hi\r\n\r\nbody")))
			})

		})
	})
}
