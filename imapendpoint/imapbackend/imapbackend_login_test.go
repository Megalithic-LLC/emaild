package imapbackend_test

import (
	"testing"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestImapBackendLogin(t *testing.T) {
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

	g.Describe("ImapBackend:Login()", func() {
		g.Before(func() {
			genjiEngine = newGenjiEngine()
			db = newDB(genjiEngine)
			accountsDAO = dao.NewAccountsDAO(db)
			mailboxesDAO = dao.NewMailboxesDAO(db)
			mailboxMessagesDAO = dao.NewMailboxMessagesDAO(db)
			messageRawBodiesDAO = dao.NewMessageRawBodiesDAO(db)
			messagesDAO = dao.NewMessagesDAO(db)
			imapBackend = newImapBackend(accountsDAO, db, mailboxesDAO, mailboxMessagesDAO, messageRawBodiesDAO, messagesDAO)
		})
		g.After(func() {
			closeAndDestroyGenjiEngine(genjiEngine)
		})

		g.It("Should refuse access to an unknown account", func() {
			_, err := imapBackend.Login(nil, "nobody", "password")
			Expect(err).Should(HaveOccurred())
		})

		g.It("Should allow access to a known account", func() {
			// setup precondition
			account := &model.Account{
				Username: "test",
				Email:    "test@acme.org",
			}
			Expect(accountsDAO.Create(account)).Should(Succeed())

			// perform test
			_, err := imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
