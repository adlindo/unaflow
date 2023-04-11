package components

import (
	"fmt"

	"github.com/adlindo/unaflow/engine"
	"github.com/maja42/goval"
)

/*
{
	"component" : "condition",
	"id" : "stepId",
	"left" : 123,
	"top" : 123,
	"data" : {
		"expr" : "(1+2) < 5",
		"trueNext" : "trueStepId",
		"falseNext" : "falseStepId"
	}
}
*/

type Condition struct {
	BaseComponent
}

func (o *Condition) Next(data map[string]interface{}, flow *engine.Flow, instance *engine.FlowInstance) (string, error) {

	eval := goval.NewEvaluator()
	variables := map[string]interface{}{}

	for name, val := range instance.ListData() {
		variables[name] = val
	}

	result, err := eval.Evaluate(o.GetStr("expr", data), variables, nil)

	if err != nil {

		fmt.Println("Error evaluate : " + o.GetStr("expr", data))
		return o.GetStr("falseNext", data), nil
	}

	if result == true {
		return o.GetStr("trueNext", data), nil
	}

	return o.GetStr("falseNext", data), nil
}
