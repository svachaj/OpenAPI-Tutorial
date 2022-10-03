package models

type System struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	ParentSystemCode string `json:"parentSystemCode,omitempty"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type Maintenance struct {
	SystemName string `json:"systemName"`
	When       string `json:"when"`
	Username   string `json:"username"`
}

type Configuration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TimeValueLog struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
