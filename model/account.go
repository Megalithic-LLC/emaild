package model

type Account struct {
	Id       string `genji:"pk"`
	Username string `genji:"index(unique)"`
	Password []byte
}
