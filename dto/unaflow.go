package dto

type Flow struct {
	Id   string
	Code string
	Name string
}

type Instance struct {
	Id     string
	FlowId string
}

type CreateInstaceReq struct {
	Execute bool
	Params  map[string]interface{}
}
