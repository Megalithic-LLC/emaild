package model

type MessageRawBody struct {
	ID   string `genji:"pk"`
	Body []byte
}
