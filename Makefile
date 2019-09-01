.PHONY: compile
compile:
	go install github.com/drauschenbach/megalithicd/cmd/megalithicd

dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
	go get -u github.com/asdine/genji/...

generate:
	genji -f model/mailbox.go  -s Mailbox
	genji -f model/property.go -s Property
