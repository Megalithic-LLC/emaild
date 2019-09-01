package main

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/model"
)

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
