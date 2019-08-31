.PHONY: compile
compile:
	go install github.com/drauschenbach/megalithicd/cmd/megalithicd

dependencies:
	dep ensure -v

generate:
	go get -u github.com/asdine/genji/...
	genji -f model/property.go -s Property
