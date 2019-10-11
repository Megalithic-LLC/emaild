package model

type Snapshot struct {
	Id        string `genji:"pk"`
	ServiceId string
	Name      string
	Engine    string
	Progress  float32
	Size      uint64
}
