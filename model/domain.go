package model

type Domain struct {
	Id                string `genji:"pk"`
	ServiceInstanceId string
	Name              string `genji:"index(unique)"`
}
