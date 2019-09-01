package model

type Message struct {
	ID       uint64 `genji:"pk"`
	DateUTC  int64  `genji:"index"`
	FlagsCSV string
}
