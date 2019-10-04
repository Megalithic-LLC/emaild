package model

type Account struct {
	Id       string `genji:"pk"`
	Email    string `genji:"index(unique)"`
	Username string `genji:"index(unique)"`
	Password []byte
}
