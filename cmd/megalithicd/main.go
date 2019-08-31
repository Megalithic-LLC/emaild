package main

import (
	"os"
	"path"

	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
	"github.com/karlkfi/inject"
)

var (
	graph       inject.Graph
	genjiEngine *engine.Engine
	db          *genji.DB
)

func init() {
	logger.SetLevel(logger.LevelDebug)
}

func main() {
	graph = inject.NewGraph()
	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newEngine))
	graph.ResolveAll()

	// If a unique id does not yet exist, create it now
	nodeid, err := getOrGenerateNodeID()
	if err != nil {
		logger.Fatalf("Failed assigning node id: %v", err)
	}

	logger.Info("Megalithic Unified Messaging started")
	logger.Infof("Node id is %s", nodeid)
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
		if _, err := tx.InitTable("properties", new(model.Property)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		logger.Fatalf("Failed initializing indexes: %v", err)
		return nil
	}

	return db
}

func newEngine() *engine.Engine {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("Failed creating current dir: %v", err)
		return nil
	}
	var eng engine.Engine
	filepath := path.Join(dir, "megalithicd.db")
	eng, err = bolt.NewEngine(filepath, 0600, nil)
	if err != nil {
		logger.Fatalf("Failed creating DB engine: %v", err)
		return nil
	}
	logger.Infof("Opened database %s", filepath)
	return &eng
}
