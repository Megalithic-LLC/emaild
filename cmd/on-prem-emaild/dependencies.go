package main

import (
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	imap_backend "github.com/emersion/go-imap/backend"
	"github.com/karlkfi/inject"
)

var (
	graph inject.Graph

	cloudService *cloudservice.CloudService
	db           *genji.DB
	genjiEngine  *engine.Engine
	imapBackend  imap_backend.Backend
	imapEndpoint *imapendpoint.ImapEndpoint

	accountsDAO         dao.AccountsDAO
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
	propertiesDAO       dao.PropertiesDAO
)

func DefineDependencies() {
	graph = inject.NewGraph()

	graph.Define(&cloudService, inject.NewAutoProvider(cloudservice.New))
	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newGenjiEngine))
	graph.Define(&imapBackend, inject.NewAutoProvider(imapbackend.New))
	graph.Define(&imapEndpoint, inject.NewAutoProvider(imapendpoint.New))

	graph.Define(&accountsDAO, inject.NewAutoProvider(dao.NewAccountsDAO))
	graph.Define(&mailboxesDAO, inject.NewAutoProvider(dao.NewMailboxesDAO))
	graph.Define(&mailboxMessagesDAO, inject.NewAutoProvider(dao.NewMailboxMessagesDAO))
	graph.Define(&messageRawBodiesDAO, inject.NewAutoProvider(dao.NewMessageRawBodiesDAO))
	graph.Define(&messagesDAO, inject.NewAutoProvider(dao.NewMessagesDAO))
	graph.Define(&propertiesDAO, inject.NewAutoProvider(dao.NewPropertiesDAO))
}
