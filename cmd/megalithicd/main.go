package main

import (
	"os"
	"path"

	"github.com/asdine/genji/engine"
	"github.com/asdine/genji/engine/bolt"
	"github.com/docktermj/go-logger/logger"
	"github.com/karlkfi/inject"
)

var (
	graph       inject.Graph
	genjiEngine engine.Engine
)

func init() {
	logger.SetLevel(logger.LevelInfo)
}

func main() {
	graph = inject.NewGraph()
	graph.Define(&genjiEngine, inject.NewAutoProvider(newEngine))
	graph.ResolveAll()
	logger.Info("Megalithicd started")
}

func newEngine() engine.Engine {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("Failed creating current dir: %v", err)
		return nil
	}
	eng, err := bolt.NewEngine(path.Join(dir, "megalithicd.db"), 0600, nil)
	if err != nil {
		logger.Fatalf("Failed creating DB engine: %v", err)
		return nil
	}
	logger.Infof("Opened database %s", eng.DB.Path())
	return eng
}
