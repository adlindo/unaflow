package components

import "github.com/adlindo/unaflow/engine"

type Join struct {
}

func (o *Join) IsAutoNext() bool {
	return true
}

func (o *Join) Execute(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) error {

	return nil
}

func (o *Join) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return "", nil
}
