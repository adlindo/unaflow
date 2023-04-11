package components

import "github.com/adlindo/unaflow/engine"

type Switch struct {
	BaseComponent
}

func (o *Switch) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	return "", nil
}
