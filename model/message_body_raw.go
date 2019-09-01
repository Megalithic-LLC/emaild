package model

type MessageBodyRaw struct {
	ID   uint64 `genji:"pk"`
	Body []byte
}
