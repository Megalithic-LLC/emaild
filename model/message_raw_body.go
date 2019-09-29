package model

type MessageRawBody struct {
	Id   string `genji:"pk"`
	Body []byte
}
