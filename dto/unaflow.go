package dto

type Flow struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Instance struct {
	Id        string                 `json:"id"`
	FlowId    string                 `json:"floWid"`
	FlowName  string                 `json:"flowName"`
	StepId    string                 `json:"stepId"`
	StepName  string                 `json:"stepName"`
	Component string                 `json:"component"`
	Status    int                    `json:"status"`
	Key       string                 `json:"key"`
	Data      map[string]interface{} `json:"data"`
}

type CreateInstaceReq struct {
	Execute bool                   `json:"execute"`
	Key     string                 `json:"key"`
	Params  map[string]interface{} `json:"params"`
}

type NextReq struct {
	Params map[string]interface{} `json:"params"`
}

type StopReq struct {
	Reason string `json:"reason"`
}
