package main

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapbackend"
	imap_backend "github.com/emersion/go-imap/backend"
	imap_server "github.com/emersion/go-imap/server"
	"github.com/karlkfi/inject"
)

var (
	graph inject.Graph

	cloudService *cloudservice.CloudService
	db           *genji.DB
	genjiEngine  *engine.Engine
	imapBackend  imap_backend.Backend
	imapServer   *imap_server.Server

	propertiesDAO dao.PropertiesDAO
)

func DefineDependencies() {
	graph = inject.NewGraph()

	graph.Define(&cloudService, inject.NewAutoProvider(cloudservice.New))
	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newGenjiEngine))
	graph.Define(&imapBackend, inject.NewAutoProvider(imapbackend.New))
	graph.Define(&imapServer, inject.NewAutoProvider(newImapServer))

	graph.Define(&propertiesDAO, inject.NewAutoProvider(dao.NewPropertiesDAO))
}
