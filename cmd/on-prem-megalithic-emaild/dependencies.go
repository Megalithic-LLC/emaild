package main

import (
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	imap_backend "github.com/emersion/go-imap/backend"
	"github.com/karlkfi/inject"
	"github.com/Megalithic-LLC/on-prem-emaild/cloudservice"
	"github.com/Megalithic-LLC/on-prem-emaild/dao"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint"
	"github.com/Megalithic-LLC/on-prem-emaild/imapendpoint/imapbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/localdelivery"
	"github.com/Megalithic-LLC/on-prem-emaild/smtpendpoint"
	"github.com/Megalithic-LLC/on-prem-emaild/smtpendpoint/smtpbackend"
	"github.com/Megalithic-LLC/on-prem-emaild/snapshotmanager"
	"github.com/Megalithic-LLC/on-prem-emaild/submissionendpoint"
	"github.com/Megalithic-LLC/on-prem-emaild/submissionendpoint/submissionbackend"
)

var (
	graph inject.Graph

	cloudService       *cloudservice.CloudService
	db                 *genji.DB
	genjiEngine        engine.Engine
	imapBackend        imap_backend.Backend
	imapEndpoint       *imapendpoint.ImapEndpoint
	localDelivery      *localdelivery.LocalDelivery
	smtpBackend        *smtpbackend.SmtpBackend
	smtpEndpoint       *smtpendpoint.SmtpEndpoint
	snapshotManager    *snapshotmanager.SnapshotManager
	submissionBackend  *submissionbackend.SubmissionBackend
	submissionEndpoint *submissionendpoint.SubmissionEndpoint

	accountsDAO         dao.AccountsDAO
	domainsDAO          dao.DomainsDAO
	endpointsDAO        dao.EndpointsDAO
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
	propertiesDAO       dao.PropertiesDAO
	snapshotsDAO        dao.SnapshotsDAO
)

func DefineDependencies() {
	graph = inject.NewGraph()

	graph.Define(&cloudService, inject.NewAutoProvider(cloudservice.New))
	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newGenjiEngine))
	graph.Define(&imapBackend, inject.NewAutoProvider(imapbackend.New))
	graph.Define(&imapEndpoint, inject.NewAutoProvider(imapendpoint.New))
	graph.Define(&localDelivery, inject.NewAutoProvider(localdelivery.New))
	graph.Define(&smtpBackend, inject.NewAutoProvider(smtpbackend.New))
	graph.Define(&smtpEndpoint, inject.NewAutoProvider(smtpendpoint.New))
	graph.Define(&snapshotManager, inject.NewAutoProvider(snapshotmanager.New))
	graph.Define(&submissionBackend, inject.NewAutoProvider(submissionbackend.New))
	graph.Define(&submissionEndpoint, inject.NewAutoProvider(submissionendpoint.New))

	graph.Define(&accountsDAO, inject.NewAutoProvider(dao.NewAccountsDAO))
	graph.Define(&domainsDAO, inject.NewAutoProvider(dao.NewDomainsDAO))
	graph.Define(&endpointsDAO, inject.NewAutoProvider(dao.NewEndpointsDAO))
	graph.Define(&mailboxesDAO, inject.NewAutoProvider(dao.NewMailboxesDAO))
	graph.Define(&mailboxMessagesDAO, inject.NewAutoProvider(dao.NewMailboxMessagesDAO))
	graph.Define(&messageRawBodiesDAO, inject.NewAutoProvider(dao.NewMessageRawBodiesDAO))
	graph.Define(&messagesDAO, inject.NewAutoProvider(dao.NewMessagesDAO))
	graph.Define(&propertiesDAO, inject.NewAutoProvider(dao.NewPropertiesDAO))
	graph.Define(&snapshotsDAO, inject.NewAutoProvider(dao.NewSnapshotsDAO))
}
