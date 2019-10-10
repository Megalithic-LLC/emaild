package model

type Snapshot struct {
	Id       string `genji:"pk"`
	Name     string
	Engine   string
	Progress float32
	Size     uint64
}
