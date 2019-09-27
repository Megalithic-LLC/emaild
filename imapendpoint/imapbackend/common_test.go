package imapbackend_test

import (
	"io/ioutil"
	"os"

	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
	. "github.com/onsi/gomega"
)

func closeAndDestroyGenjiEngine(engine *engine.Engine) {
	engine_ := *engine
	boltEngine, ok := engine_.(*bolt.Engine)
	Expect(ok).To(Equal(true))
	boltDB := boltEngine.DB
	dbPath := boltDB.Path()
	Expect(boltDB.Close()).Should(Succeed())
	Expect(os.Remove(dbPath)).Should(Succeed())
}

func newDB(engine *engine.Engine) *genji.DB {
	db, err := genji.New(*engine)
	if err != nil {
		logger.Fatalf("Failed creating database engine: %v", err)
		return nil
	}

	// Initialize tables, creating indexes when needed
	logger.Debugf("Ensuring indexes")
	if err := db.Update(func(tx *genji.Tx) error {
		for tableName, tableModel := range model.Tables {
			if _, err := tx.InitTable(tableName, tableModel); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Fatalf("Failed initializing indexes: %v", err)
		return nil
	}

	return db
}

func newGenjiEngine() *engine.Engine {
	file, err := ioutil.TempFile("", "on-prem-emaild-testing-")
	if err != nil {
		logger.Fatalf("Failed creating DB for testing: %v", err)
		return nil
	}
	file.Close()
	os.Remove(file.Name())
	var eng engine.Engine
	eng, err = bolt.NewEngine(file.Name(), 0600, nil)
	if err != nil {
		logger.Fatalf("Failed creating DB engine: %v", err)
		return nil
	}
	logger.Debugf("Opened database %s", file.Name())
	return &eng
}

func newImapBackend(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	mailboxesDAO dao.MailboxesDAO,
) *imapbackend.ImapBackend {
	return imapbackend.New(accountsDAO, db, mailboxesDAO)
}
