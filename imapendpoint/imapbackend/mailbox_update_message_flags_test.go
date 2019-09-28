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
	"github.com/emersion/go-imap"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestMailboxUpdateMessageFlags(t *testing.T) {
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

	g.Describe("Mailbox flag update operations", func() {
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

		g.It("Should correctly set a new flag additively", func() {
			// setup precondition
			account := &model.Account{Username: "test"}
			Expect(accountsDAO.Create(account)).Should(Succeed())
			user, err := imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
			mailbox, err := user.GetMailbox("INBOX")
			Expect(err).ToNot(HaveOccurred())
			Expect(mailbox).ToNot(BeNil())
			{
				flags := []string{"\\Flagged"}
				var date time.Time
				body := "Subject: hi\r\n\r\nbody"
				Expect(mailbox.CreateMessage(flags, date, strings.NewReader(body))).ToNot(HaveOccurred())
			}

			// Sanity check - fetch the message and expect just the \Flagged flag to be set
			messages := make(chan *imap.Message, 1)
			uid := false
			seqSet := new(imap.SeqSet)
			seqSet.AddRange(1, 1)
			items := []imap.FetchItem{imap.FetchItem(imap.FetchFlags)}
			Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
			message := <-messages
			Expect(message.SeqNum).To(Equal(uint32(1)))
			flagsUntyped := message.Items[imap.FetchFlags]
			Expect(flagsUntyped).ToNot(BeNil())
			flags := flagsUntyped.([]string)
			Expect(flags).To(ContainElement("\\Flagged"))
			Expect(flags).To(HaveLen(1))

			// Perform test
			Expect(mailbox.UpdateMessagesFlags(uid, seqSet, imap.AddFlags, []string{"\\Recent"})).ToNot(HaveOccurred())
			Expect(mailbox.ListMessages(uid, seqSet, items, messages)).ToNot(HaveOccurred())
			message = <-messages
			flagsUntyped = message.Items[imap.FetchFlags]
			Expect(flagsUntyped).ToNot(BeNil())
			flags = flagsUntyped.([]string)
			Expect(flags).To(ContainElement("\\Flagged"))
			Expect(flags).To(ContainElement("\\Recent"))
			Expect(flags).To(HaveLen(2))
		})
	})
}
