package model

type MessageBodyRaw struct {
	ID     uint32 `genji:"pk"`
	Body   []byte
	Unused uint32 `genji:"index"`
}
