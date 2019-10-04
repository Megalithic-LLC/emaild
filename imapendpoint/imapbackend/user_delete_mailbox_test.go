package imapbackend_test

import (
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

func TestUserDeleteMailbox(t *testing.T) {
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

	g.Describe("User", func() {

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

		g.Describe("DeleteMailbox()", func() {

			g.It("Should remove cross-post information", func() {

				// Setup precondition
				accountModel := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(accountModel)).Should(Succeed())
				mailboxModel := &model.Mailbox{AccountId: accountModel.Id, Name: "foo"}
				Expect(mailboxesDAO.Create(mailboxModel)).Should(Succeed())
				time1 := time.Date(2010, 2, 20, 22, 7, 15, 0, time.UTC)
				messageModel := &model.Message{DateUTC: time1.UTC().Unix()}
				Expect(messagesDAO.Create(messageModel)).Should(Succeed())
				mailboxMessageModel := &model.MailboxMessage{MailboxId: mailboxModel.Id, MessageId: messageModel.Id}
				Expect(mailboxMessagesDAO.Create(mailboxMessageModel)).Should(Succeed())

				// Verify precondition
				user, err := imapBackend.Login(nil, "test", "password")
				mailbox, err := user.GetMailbox("foo")
				Expect(err).ToNot(HaveOccurred())
				messages := make(chan *imap.Message, 1)
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 1)
				items := []imap.FetchItem{imap.FetchInternalDate}
				Expect(mailbox.ListMessages(false, seqSet, items, messages)).To(Succeed())
				message1 := <-messages
				Expect(message1.InternalDate.UTC()).To(Equal(time1.UTC()))

				// Perform test
				Expect(user.DeleteMailbox("foo")).To(Succeed())
				mailboxMessageModel_, err := mailboxMessagesDAO.FindByIds(mailboxModel.Id, messageModel.Id)
				Expect(err).ToNot(HaveOccurred())
				Expect(mailboxMessageModel_).To(BeNil())
			})

			g.It("Should garbage collect a message no longer cross-posted into any mailboxes", func() {

				// Setup precondition
				accountModel := &model.Account{Username: "test"}
				Expect(accountsDAO.Create(accountModel)).Should(Succeed())
				mailboxModel := &model.Mailbox{AccountId: accountModel.Id, Name: "foo"}
				Expect(mailboxesDAO.Create(mailboxModel)).Should(Succeed())
				time1 := time.Date(2010, 2, 20, 22, 7, 15, 0, time.UTC)
				messageModel := &model.Message{DateUTC: time1.UTC().Unix()}
				Expect(messagesDAO.Create(messageModel)).Should(Succeed())
				mailboxMessageModel := &model.MailboxMessage{MailboxId: mailboxModel.Id, MessageId: messageModel.Id}
				Expect(mailboxMessagesDAO.Create(mailboxMessageModel)).Should(Succeed())

				// Verify precondition
				user, err := imapBackend.Login(nil, "test", "password")
				mailbox, err := user.GetMailbox("foo")
				Expect(err).ToNot(HaveOccurred())
				messages := make(chan *imap.Message, 1)
				seqSet := new(imap.SeqSet)
				seqSet.AddRange(1, 1)
				items := []imap.FetchItem{imap.FetchInternalDate}
				Expect(mailbox.ListMessages(false, seqSet, items, messages)).To(Succeed())
				message1 := <-messages
				Expect(message1.InternalDate.UTC()).To(Equal(time1.UTC()))

				// Perform test
				Expect(user.DeleteMailbox("foo")).To(Succeed())
				messageModel_, err := messagesDAO.FindById(messageModel.Id)
				Expect(err).ToNot(HaveOccurred())
				Expect(messageModel_).To(BeNil())
			})

		})
	})
}
