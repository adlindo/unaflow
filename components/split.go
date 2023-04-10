package components

import "github.com/adlindo/unaflow/engine"

type Split struct {
}

func (o *Split) IsAutoNext() bool {
	return true
}

func (o *Split) Execute(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) error {

	return nil
}

func (o *Split) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return "", nil
}
