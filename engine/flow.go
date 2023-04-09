package engine

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/adlindo/unaflow/repo"
)

type Flow struct {
	Id    string
	Steps map[string]*Step
}

/*
	{
		"name" : "flow name",
		"steps" : [
			{
				"component" : "componentCode",
				"id" : "stepId",
				"left" : 123,
				"top" : 123,
				"data" : {
					"param1" : "string",
					"param2" : 123,
					"param3" : 123.456,
					"param4" : true,
					"next" : "nextStepId"
				}
			}
		]
	}
*/

func (o *Flow) Init(mdl *repo.Flow) error {

	o.Steps = map[string]*Step{}

	scriptMap := map[string]interface{}{}

	err := json.Unmarshal([]byte(mdl.Script), &scriptMap)

	if err != nil {

		return errors.New("unable to parse script")
	}

	steps := scriptMap["steps"].([]interface{})

	for _, stepItem := range steps {

		fmt.Println("===>>>33333", stepItem)

		stepMap := stepItem.(map[string]interface{})

		fmt.Println("===>>>444444", stepMap, "========>", stepMap["component"].(string))

		if IsCompExist(stepMap["component"].(string)) {

			fmt.Println("===>>>5555555", stepMap["data"].(map[string]interface{}))

			o.Steps[stepMap["id"].(string)] = &Step{
				Flow: o,
				Data: stepMap["data"].(map[string]interface{}),
			}
		} else {
			fmt.Println("===>>>66666666", stepItem)
		}
	}

	return nil
}

func (o *Flow) Start(params map[string]interface{}, autoExecute bool) *FlowInstance {

	ret := &FlowInstance{}

	return ret
}

func (o *Flow) GetStep(stepId string) *Step {

	ret, ok := o.Steps[stepId]

	if !ok {

	}

	return ret
}

func (o *Flow) Execute(instance *FlowInstance) error {

	prevStepId := instance.GetStepId()

	step := o.GetStep(prevStepId)

	if step == nil {
		return errors.New("step not found :" + instance.GetStepId())
	}

	next, err := step.Execute(instance)

	if err != nil {

		return err
	}

	if next != prevStepId {

		instance.SetStepId(next)
	}

	return nil
}

// -------------------------

var flowMap map[string]*Flow = map[string]*Flow{}

func GetFlow(flowId string) (*Flow, error) {

	ret, ok := flowMap[flowId]

	if !ok {

		mdl := repo.GetFlowRepo().GetById(flowId)

		if mdl == nil {

			return nil, errors.New("flow not found : " + flowId)
		}

		ret = &Flow{
			Id: flowId,
		}

		err := ret.Init(mdl)

		if err != nil {
			return nil, err
		}

		flowMap[flowId] = ret
	}

	return ret, nil
}
