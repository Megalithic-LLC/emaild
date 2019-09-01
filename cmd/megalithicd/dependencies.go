package main

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/drauschenbach/megalithicd/imapbackend"
	imap_backend "github.com/emersion/go-imap/backend"
	imap_server "github.com/emersion/go-imap/server"
	"github.com/karlkfi/inject"
)

var (
	graph inject.Graph

	db          *genji.DB
	genjiEngine *engine.Engine
	imapBackend imap_backend.Backend
	imapServer  *imap_server.Server
)

func DefineDependencies() {
	graph = inject.NewGraph()

	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newGenjiEngine))
	graph.Define(&imapBackend, inject.NewAutoProvider(imapbackend.New))
	graph.Define(&imapServer, inject.NewAutoProvider(newImapServer))
}
