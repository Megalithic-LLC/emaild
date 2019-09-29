package model

type Mailbox struct {
	Id          string `genji:"pk"`
	AccountId   string `genji:"index"`
	Name        string `genji:"index"`
	Messages    uint32
	Recent      uint32
	Unseen      uint32
	UidNext     uint32
	UidValidity uint32
	Subscribed  bool
}
