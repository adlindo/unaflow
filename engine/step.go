package engine

import "errors"

type Step struct {
	Flow      *Flow
	Component string
	Id        string
	Name      string
	Data      map[string]interface{}
}

func (o *Step) Execute(instance *FlowInstance) error {

	comp := GetComp(o.Component)

	if comp == nil {
		return errors.New("component not found :" + o.Component)
	}

	return comp.Execute(o.Data, o.Flow, instance)
}

func (o *Step) Next(instance *FlowInstance) (string, error) {

	comp := GetComp(o.Component)

	if comp == nil {
		return "", errors.New("component not found :" + o.Component)
	}

	return comp.Next(o.Data, o.Flow, instance)
}

func (o *Step) IsAutoNext() bool {

	comp := GetComp(o.Component)

	if comp != nil {
		return comp.IsAutoNext()
	}

	return false
}
