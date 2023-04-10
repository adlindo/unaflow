package components

import "github.com/adlindo/unaflow/engine"

type Condition struct {
	BaseComponent
}

func (o *Condition) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return "", nil
}
