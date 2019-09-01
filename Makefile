.PHONY: compile
compile:
	go install github.com/drauschenbach/megalithicd/cmd/megalithicd

dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
	go get -u github.com/asdine/genji/...

generate:
	genji -f model/mailbox.go          -s Mailbox
	genji -f model/mailbox_message.go  -s MailboxMessage
	genji -f model/message.go          -s Message
	genji -f model/message_body_raw.go -s MessageBodyRaw
	genji -f model/property.go         -s Property
