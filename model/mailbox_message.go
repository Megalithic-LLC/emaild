package model

type MailboxMessage struct {
	MailboxID string `genji:"index"`
	MessageID uint32 `genji:"index"`
}
