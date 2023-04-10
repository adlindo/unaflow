package engine

import (
	"errors"

	"github.com/adlindo/unaflow/repo"
)

type FlowInstance struct {
	Id  string
	mdl *repo.Instance
}

func (o *FlowInstance) Execute(param map[string]interface{}) error {

	o.mdl = repo.GetInstanceRepo().GetById(o.Id)

	if o.mdl == nil {
		return errors.New("instance not found : " + o.Id)
	}

	flow, err := GetFlow(o.mdl.FlowId)

	if err != nil {
		return errors.New("flow not found :" + o.mdl.FlowId)
	}

	if param != nil {

		o.SetDataBulk(param)
	}

	return flow.Execute(o)
}

func (o *FlowInstance) Next(param map[string]interface{}) error {

	o.mdl = repo.GetInstanceRepo().GetById(o.Id)

	if o.mdl == nil {
		return errors.New("instance not found : " + o.Id)
	}

	flow, err := GetFlow(o.mdl.FlowId)

	if err != nil {
		return errors.New("flow not found :" + o.mdl.FlowId)
	}

	if param != nil {

		o.SetDataBulk(param)
	}

	return flow.Next(o)
}

func (o *FlowInstance) GetStepId() string {

	if o.mdl == nil {
		return ""
	}

	return o.mdl.StepId
}

func (o *FlowInstance) SetStepId(stepId string) {

	repo.GetInstanceRepo().SetStepId(o.Id, stepId)
}

func (o *FlowInstance) SetData(name string, val interface{}) {

	repo.GetInstanceDataRepo().SetData(o.Id, name, val)
}

func (o *FlowInstance) SetDataBulk(data map[string]interface{}) {

	for name, val := range data {
		o.SetData(name, val)
	}
}

func (o *FlowInstance) GetData(name string) interface{} {

	return repo.GetInstanceDataRepo().GetData(o.Id, name)
}

func (o *FlowInstance) ListData() map[string]interface{} {

	return repo.GetInstanceDataRepo().ListData(o.Id)
}

//--------------------------------------------------

func GetInstace(instanceId string) *FlowInstance {

	mdl := repo.GetInstanceRepo().GetById(instanceId)

	if mdl == nil {

		return nil
	}

	return &FlowInstance{
		Id: instanceId,
	}
}
