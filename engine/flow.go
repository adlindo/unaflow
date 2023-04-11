package engine

import (
	"encoding/json"
	"errors"

	"github.com/adlindo/unaflow/dto"
	"github.com/adlindo/unaflow/repo"
	"github.com/jinzhu/copier"
)

type Flow struct {
	Id     string
	Name   string
	Code   string
	Active bool
	Steps  map[string]*Step
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

		stepMap := stepItem.(map[string]interface{})

		if IsCompExist(stepMap["component"].(string)) {

			o.Steps[stepMap["id"].(string)] = &Step{
				Flow:      o,
				Id:        stepMap["id"].(string),
				Name:      stepMap["name"].(string),
				Component: stepMap["component"].(string),
				Data:      stepMap["data"].(map[string]interface{}),
			}
		}
	}

	return nil
}

func (o *Flow) CreateInstance(key string, data map[string]interface{}, execute bool) (*FlowInstance, error) {

	if !o.Active {

		return nil, errors.New("Flow is not active")
	}

	mdl := &repo.Instance{
		FlowId: o.Id,
		Key:    key,
		StepId: "start",
		Status: 0,
	}

	repo.GetInstanceRepo().Create(mdl)

	ret := &FlowInstance{
		Id:       mdl.Id,
		FlowId:   o.Id,
		FlowName: o.Name,
		StepId:   "start",
		Status:   0,
	}

	if data != nil {
		ret.SetDataBulk(data)
	}

	if execute {
		ret.SetStatus(1)
		ret.Execute(nil)
	} else {

		step := o.getStep("start")

		if step != nil {

			ret.Component = step.Component
			ret.StepName = step.Name
		}
	}

	return ret, nil
}

func (o *Flow) getStep(stepId string) *Step {

	ret, ok := o.Steps[stepId]

	if !ok {

		return nil
	}

	return ret
}

func (o *Flow) Execute(instance *FlowInstance) error {

	step := o.getStep(instance.GetStepId())

	if step == nil {
		return errors.New("step not found :" + instance.GetStepId())
	}

	instance.FlowId = o.Id
	instance.FlowName = o.Name
	instance.StepId = step.Id
	instance.StepName = step.Name
	instance.Component = step.Component

	err := step.Execute(instance)

	if err != nil {

		return err
	}

	if step.Component == "end" {

		instance.End()
	} else if step.IsAutoNext() {

		instance.Next(nil)
	}

	return nil
}

func (o *Flow) Next(instance *FlowInstance) error {

	prevStepId := instance.GetStepId()

	step := o.getStep(prevStepId)

	if step == nil {
		return errors.New("step not found :" + instance.GetStepId())
	}

	next, err := step.Next(instance)

	if err != nil {

		return err
	}

	if next != prevStepId {

		instance.SetStepId(next)
		instance.Execute(nil)
	}

	return nil
}

func (o *Flow) ListInstance(step, filter string, pageNo, pageSize int) ([]dto.Instance, int) {

	ret := []dto.Instance{}
	list, total := repo.GetInstanceRepo().Search(o.Id, step, filter, pageNo, pageSize)

	for _, item := range list {

		newDto := dto.Instance{}
		copier.Copy(&newDto, item)

		newDto.FlowName = o.Name

		step := o.getStep(newDto.StepId)

		if step != nil {
			newDto.StepName = step.Name
			newDto.Component = step.Component
		}

		ret = append(ret, newDto)
	}

	return ret, int(total)
}

// -------------------------

var flowMap map[string]*Flow = map[string]*Flow{}

func ClearFlow() {
	flowMap = map[string]*Flow{}
}

func ListFlow(filter string, pageNo, pageSize int) ([]dto.Flow, int64) {

	ret := []dto.Flow{}

	listMdl, total := repo.GetFlowRepo().Search(filter, pageNo, pageSize)

	for _, mdl := range listMdl {

		newDto := dto.Flow{}

		copier.Copy(&newDto, mdl)
		ret = append(ret, newDto)
	}

	return ret, total
}

func GetFlow(idOrCode string) (*Flow, error) {

	ret, ok := flowMap[idOrCode]

	if !ok {

		mdl := repo.GetFlowRepo().GetById(idOrCode)

		if mdl == nil {

			mdl = repo.GetFlowRepo().GetByCode(idOrCode)

			if mdl == nil {

				return nil, errors.New("flow not found : " + idOrCode)
			}
		}

		ret = &Flow{
			Id:     mdl.Id,
			Name:   mdl.Name,
			Code:   mdl.Code,
			Active: mdl.IsActive,
		}

		err := ret.Init(mdl)

		if err != nil {
			return nil, err
		}

		flowMap[mdl.Id] = ret
		flowMap[mdl.Code] = ret
	}

	return ret, nil
}
