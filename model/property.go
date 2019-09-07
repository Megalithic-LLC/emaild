package model

type Property struct {
	Key    string `genji:"pk"`
	Value  string
	Unused uint32 `genji:"index"`
}
