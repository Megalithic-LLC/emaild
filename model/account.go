package model

type Account struct {
	Id          string `genji:"pk"`
	Name        string
	Email       string `genji:"index(unique)"`
	First       string
	Last        string
	DisplayName string
	Password    []byte
}
