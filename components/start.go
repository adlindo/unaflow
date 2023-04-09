package components

import "github.com/adlindo/unaflow/engine"

type Start struct {
}

func (o *Start) IsAutoExecute() bool {
	return true
}

func (o *Start) Execute() (string, error) {

	return "", nil
}

func init() {
	engine.RegComponent("start", &Start{})
	engine.RegComponent("end", &End{})
	engine.RegComponent("userTask", &UserTask{})
	engine.RegComponent("split", &Split{})
	engine.RegComponent("join", &Join{})
	engine.RegComponent("condition", &Condition{})
	engine.RegComponent("callHTTP", &CallHTTP{})
}
