package model

type Mailbox struct {
	ID          string `genji:"pk"`
	Name        string `genji:"index"`
	Messages    uint32
	Recent      uint32
	Unseen      uint32
	UidNext     uint32
	UidValidity uint32
}
