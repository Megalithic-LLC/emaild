package submissionbackend_test

import (
	"strings"
	"testing"

	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/franela/goblin"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/localdelivery"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/Megalithic-LLC/on-prem-emaild/submissionendpoint/submissionbackend"
	. "github.com/onsi/gomega"
)

func TestLocalDelivery(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	var genjiEngine *engine.Engine
	var db *genji.DB
	var submissionBackend *submissionbackend.SubmissionBackend
	var accountsDAO dao.AccountsDAO
	var localDelivery *localdelivery.LocalDelivery
	var mailboxesDAO dao.MailboxesDAO
	var mailboxMessagesDAO dao.MailboxMessagesDAO
	var messageRawBodiesDAO dao.MessageRawBodiesDAO
	var messagesDAO dao.MessagesDAO

	g.Describe("submissionBackend", func() {
		g.BeforeEach(func() {
			genjiEngine = newGenjiEngine()
			db = newDB(genjiEngine)
			accountsDAO = dao.NewAccountsDAO(db)
			mailboxesDAO = dao.NewMailboxesDAO(db)
			mailboxMessagesDAO = dao.NewMailboxMessagesDAO(db)
			messageRawBodiesDAO = dao.NewMessageRawBodiesDAO(db)
			messagesDAO = dao.NewMessagesDAO(db)
			localDelivery = localdelivery.New(accountsDAO, db, mailboxesDAO, mailboxMessagesDAO, messageRawBodiesDAO, messagesDAO)
			submissionBackend = newSubmissionBackend(accountsDAO, db, localDelivery, mailboxesDAO, mailboxMessagesDAO, messageRawBodiesDAO, messagesDAO)
		})
		g.AfterEach(func() {
			closeAndDestroyGenjiEngine(genjiEngine)
		})

		g.Describe("Local delivery", func() {

			g.It("Should correctly deliver to a single recipient", func() {
				// setup precondition
				account := &model.Account{Name: "test", Email: "test@acme.org"}
				Expect(accountsDAO.Create(account)).To(Succeed())

				// Perform delivery
				session, err := submissionBackend.Login(nil, "test@acme.org", "password")
				Expect(err).ToNot(HaveOccurred())
				Expect(session.Mail("")).To(Succeed())
				Expect(session.Rcpt("test@acme.org")).To(Succeed())
				Expect(session.Data(strings.NewReader("Subject: A\r\n\r\nbody"))).To(Succeed())

				// Verify delivery
				mailbox, err := mailboxesDAO.FindOneByName(account.Id, "INBOX")
				Expect(err).ToNot(HaveOccurred())
				where := model.NewMailboxMessageFields().MailboxId.Eq(mailbox.Id)
				var mailboxMessage *model.MailboxMessage
				err = mailboxMessagesDAO.Find(where, 1, func(mailboxMessage_ *model.MailboxMessage) error {
					mailboxMessage = mailboxMessage_
					return nil
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(mailboxMessage).ToNot(BeNil())
				messageRawBody, err := messageRawBodiesDAO.FindById(mailboxMessage.MessageId)
				Expect(err).ToNot(HaveOccurred())
				Expect(messageRawBody.Body).To(Equal([]byte("Subject: A\r\n\r\nbody")))
			})

		})
	})
}
