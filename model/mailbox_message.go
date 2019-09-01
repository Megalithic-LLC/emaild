package model

type MailboxMessage struct {
	MailboxID string `genji:"index"`
	MessageID uint64 `genji:"index"`
}
