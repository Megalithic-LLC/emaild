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

func Test(t *testing.T) {
	logger.SetLevel(logger.LevelDebug)

	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	var genjiEngine *engine.Engine
	var db *genji.DB
	var imapBackend *imapbackend.ImapBackend
	var accountsDAO dao.AccountsDAO

	g.Describe("ImapBackend:Login()", func() {
		g.Before(func() {
			genjiEngine = newGenjiEngine()
			db = newDB(genjiEngine)
			accountsDAO = dao.NewAccountsDAO(db)
			imapBackend = newImapBackend(accountsDAO, db)
		})
		g.After(func() {
			closeAndDestroyGenjiEngine(genjiEngine)
		})

		g.It("Should refuse access to an unknown account", func() {
			_, err := imapBackend.Login(nil, "nobody", "password")
			Expect(err).Should(HaveOccurred())
		})

		g.It("Should allow access to a known account", func() {
			account := model.Account{
				Username: "test",
			}
			Expect(accountsDAO.Create(&account)).Should(Succeed())
			_, err := imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
