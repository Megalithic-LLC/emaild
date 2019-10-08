package model

type Account struct {
	Id                string `genji:"pk"`
	ServiceInstanceId string
	Name              string
	DomainId          string
	Email             string `genji:"index(unique)"`
	First             string
	Last              string
	DisplayName       string
	Password          []byte
}
