package imapbackend

import (
	"github.com/asdine/genji"
)

type ImapBackend struct {
	db *genji.DB
}

func New(db *genji.DB) *ImapBackend {
	self := ImapBackend{
		db: db,
	}
	return &self
}
