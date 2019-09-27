package imapbackend_test

import (
	"testing"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/docktermj/go-logger/logger"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	logger.SetLevel(logger.LevelDebug)

	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	var genjiEngine *engine.Engine
	var db *genji.DB
	var imapBackend *imapbackend.ImapBackend
	var accountsDAO dao.AccountsDAO
	var mailboxesDAO dao.MailboxesDAO

	g.Describe("User operations", func() {
		g.BeforeEach(func() {
			genjiEngine = newGenjiEngine()
			db = newDB(genjiEngine)
			accountsDAO = dao.NewAccountsDAO(db)
			mailboxesDAO = dao.NewMailboxesDAO(db)
			imapBackend = newImapBackend(accountsDAO, db, mailboxesDAO)
		})
		g.AfterEach(func() {
			closeAndDestroyGenjiEngine(genjiEngine)
		})

		g.It("Should refuse access to a non-existent mailbox", func() {
			// setup precondition
			account := &model.Account{Username: "test"}
			Expect(accountsDAO.Create(account)).Should(Succeed())

			// Perform test
			user, err := imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
			_, err = user.GetMailbox("nonexistent")
			Expect(err).Should(HaveOccurred())
		})

		g.It("Should allow access to a known mailbox", func() {
			// setup precondition
			account := &model.Account{Username: "test"}
			Expect(accountsDAO.Create(account)).Should(Succeed())
			user, err := imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
			Expect(user.CreateMailbox("foo")).ToNot(HaveOccurred())

			// Perform test
			_, err = user.GetMailbox("foo")
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
