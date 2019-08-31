package model

type Property struct {
	ID    string `genji:"pk"`
	Name  string `genji:"index(unique)"`
	Value string
}
