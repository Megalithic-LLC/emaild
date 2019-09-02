package model

type Message struct {
	ID       uint32 `genji:"pk"`
	DateUTC  int64  `genji:"index"`
	FlagsCSV string
}
