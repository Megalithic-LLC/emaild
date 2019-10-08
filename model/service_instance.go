package model

type ServiceInstance struct {
	Id        string `genji:"pk"`
	ServiceId string
	PlanId    string
}
