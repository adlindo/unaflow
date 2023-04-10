package engine

import "errors"

type Step struct {
	Flow          *Flow
	ComponentCode string
	Data          map[string]interface{}
}

func (o *Step) Execute(instance *FlowInstance) error {

	comp := GetComp(o.ComponentCode)

	if comp == nil {
		return errors.New("component not found :" + o.ComponentCode)
	}

	return comp.Execute(o.Data, o.Flow, instance)
}

func (o *Step) Next(instance *FlowInstance) (string, error) {

	comp := GetComp(o.ComponentCode)

	if comp == nil {
		return "", errors.New("component not found :" + o.ComponentCode)
	}

	return comp.Next(o.Data, o.Flow, instance)
}

func (o *Step) IsAutoNext() bool {

	comp := GetComp(o.ComponentCode)

	if comp != nil {
		return comp.IsAutoNext()
	}

	return false
}
