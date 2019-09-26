package model

type Account struct {
	ID       uint32 `genji:"pk"`
	Username string `genji:"index(unique)"`
	Password []byte
}
