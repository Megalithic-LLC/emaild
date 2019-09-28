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
	"github.com/docktermj/go-logger/logger"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMailboxCreateMessage(t *testing.T) {
	logger.SetLevel(logger.LevelDebug)

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
				flags := []string{}
				var date time.Time
				body := `Subject: hi\r\n\r\nbody`
				Expect(mailbox.CreateMessage(flags, date, strings.NewReader(body))).ToNot(HaveOccurred())
				//TODO ListMessages(uid bool, seqSet *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
			})

		})
	})
}
