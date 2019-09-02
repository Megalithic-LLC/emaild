package model

type MessageBodyRaw struct {
	ID   uint32 `genji:"pk"`
	Body []byte
}
