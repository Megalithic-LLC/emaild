package model

type Account struct {
	ID       string `genji:"pk"`
	Username string `genji:"index(unique)"`
	Password []byte
}
