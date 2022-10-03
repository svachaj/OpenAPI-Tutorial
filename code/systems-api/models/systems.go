package models

import "time"

type System struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	ParentSystemCode string `json:"parentSystemCode"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type Maintenance struct {
	SystemName string    `json:"systemName"`
	When       time.Time `json:"when"`
	Username   string    `json:"username"`
}

type Configuration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TimeValueLog struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
	Unit  string    `json:"unit"`
}
