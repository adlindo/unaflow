package components

import "github.com/adlindo/unaflow/engine"

type MultiCondition struct {
	BaseComponent
}

func (o *MultiCondition) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return "", nil
}
