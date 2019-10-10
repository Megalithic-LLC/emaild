package model

type Endpoint struct {
	Id                string `genji:"pk"`
	ServiceInstanceId string
	Protocol          string
	Type              string
	Port              uint16
	Path              string
	Enabled           bool
}
