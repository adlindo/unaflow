package components

import "github.com/adlindo/unaflow/engine"

type BaseComponent struct {
}

func (o *BaseComponent) GetInt(name string, data map[string]interface{}) int {

	val, ok := data[name]

	if !ok {
		return 0
	}

	return val.(int)
}

func (o *BaseComponent) GetStr(name string, data map[string]interface{}) string {

	val, ok := data[name]

	if !ok {
		return ""
	}

	return val.(string)
}

func (o *BaseComponent) GetBool(name string, data map[string]interface{}) bool {

	val, ok := data[name]

	if !ok {
		return false
	}

	return val.(bool)
}

func (o *BaseComponent) GetFloat(name string, data map[string]interface{}) float64 {

	val, ok := data[name]

	if !ok {
		return 0
	}

	return val.(float64)
}

func (o *BaseComponent) Execute(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) error {

	return nil
}

func (o *BaseComponent) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return o.GetStr("next", data), nil
}

func (o *BaseComponent) IsAutoNext() bool {
	return true
}

//------------------------------------------------------------------------

func init() {
	engine.RegComponent("start", &Start{})
	engine.RegComponent("end", &End{})
	engine.RegComponent("userTask", &UserTask{})
	engine.RegComponent("split", &Split{})
	engine.RegComponent("join", &Join{})
	engine.RegComponent("condition", &Condition{})
	engine.RegComponent("switch", &Switch{})
	engine.RegComponent("callHTTP", &CallHTTP{})
	engine.RegComponent("sendEmail", &SendEmail{})
	engine.RegComponent("setValue", &SetValue{})
}
