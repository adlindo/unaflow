package engine

import "fmt"

type Component interface {
	IsAutoExecute() bool
	Execute() (string, error)
}

var compMap map[string]Component = map[string]Component{}

func RegComponent(code string, comp Component) {

	fmt.Println("register ===>>>", code)
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
