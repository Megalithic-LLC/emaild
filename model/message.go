package model

type Message struct {
	Id      string `genji:"pk"`
	DateUTC int64  `genji:"index"`
	Size    uint32
}
