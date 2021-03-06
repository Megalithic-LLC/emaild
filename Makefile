.PHONY: compile
compile:
	go install ./...

dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
	go get -u github.com/asdine/genji/...

generate:
	genji -f model/account.go          -s Account
	genji -f model/domain.go           -s Domain
	genji -f model/endpoint.go         -s Endpoint
	genji -f model/mailbox.go          -s Mailbox
	genji -f model/mailbox_message.go  -s MailboxMessage
	genji -f model/message.go          -s Message
	genji -f model/message_raw_body.go -s MessageRawBody
	genji -f model/property.go         -s Property
	genji -f model/snapshot.go         -s Snapshot
	go generate ./...

check:
	go test -v ./...
