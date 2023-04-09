package engine

type Step struct {
	Flow          *Flow
	ComponentCode string
	Data          map[string]interface{}
}

func (o *Step) Execute(instance *FlowInstance) (string, error) {

	return "", nil
}
