package model

type Message struct {
	ID      string `genji:"pk"`
	DateUTC int64  `genji:"index"`
	Size    uint32
}
