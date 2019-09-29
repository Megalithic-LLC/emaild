package model

type Mailbox struct {
	Id          string `genji:"pk"`
	AccountId   string `genji:"index"`
	Name        string `genji:"index"`
	UidNext     uint32
	UidValidity uint32
	Subscribed  bool
}
