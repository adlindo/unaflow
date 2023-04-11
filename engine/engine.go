package engine

type Component interface {
	IsAutoNext() bool
	Execute(data map[string]interface{}, flow *Flow, instance *FlowInstance) error
	Next(data map[string]interface{}, flow *Flow, instance *FlowInstance) (string, error)
}

var compMap map[string]Component = map[string]Component{}

func RegComponent(code string, comp Component) {

	compMap[code] = comp
}

func GetComp(code string) Component {

	ret, ok := compMap[code]

	if ok {

		return ret
	}

	return nil
}

func IsCompExist(code string) bool {

	_, ok := compMap[code]

	return ok
}
