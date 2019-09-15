package model

type Account struct {
	ID       string `genji:"pk"`
	Name     string
	Provider string
	Email    string `genji:"index(unique)"`

	ImapHost     string
	ImapPort     uint16
	ImapUsername string
	ImapPassword string
	SmtpHost     string
	SmtpPort     uint16
	SmtpUsername string
	SmtpPassword string
	SslRequired  bool
}
