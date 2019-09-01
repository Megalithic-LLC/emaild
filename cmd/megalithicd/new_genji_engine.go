package main

import (
	"os"
	"path"

	"github.com/asdine/genji/engine"
	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
)

func newGenjiEngine() *engine.Engine {
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
