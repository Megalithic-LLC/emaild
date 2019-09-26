package imapbackend_test

import (
	"os"
	"testing"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	boltengine "github.com/asdine/genji/engine/bolt"
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
			genjiEngine_ := *genjiEngine
			boltEngine, ok := genjiEngine_.(*boltengine.Engine)
			Expect(ok).To(Equal(true))
			boltDB := boltEngine.DB
			dbPath := boltDB.Path()
			Expect(boltDB.Close()).ToNot(HaveOccurred())
			Expect(os.Remove(dbPath)).ToNot(HaveOccurred())
		})

		g.It("Should refuse access to an unknown account", func() {
			_, err := imapBackend.Login(nil, "nobody", "password")
			Expect(err).Should(HaveOccurred())
		})

		g.It("Should allow access to a known account", func() {
			account := model.Account{
				Username: "test",
			}
			err := accountsDAO.Create(&account)
			Expect(err).ToNot(HaveOccurred())
			_, err = imapBackend.Login(nil, "test", "password")
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
