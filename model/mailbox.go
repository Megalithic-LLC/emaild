package model

type Mailbox struct {
	ID          string `genji:"pk"`
	AccountID   string `genji:"index"`
	Name        string `genji:"index"`
	Messages    uint32
	Recent      uint32
	Unseen      uint32
	UidNext     uint32
	UidValidity uint32
	Subscribed  bool
}
