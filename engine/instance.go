package engine

import (
	"errors"

	"github.com/adlindo/unaflow/repo"
)

type FlowInstance struct {
	Id        string
	StepId    string
	StepName  string
	Component string
	FlowId    string
	FlowName  string
	Status    int
	mdl       *repo.Instance
}

func (o *FlowInstance) GetFlow() (*Flow, error) {

	o.mdl = repo.GetInstanceRepo().GetById(o.Id)

	if o.mdl == nil {
		return nil, errors.New("instance not found : " + o.Id)
	}

	flow, err := GetFlow(o.mdl.FlowId)

	if err != nil {
		return nil, errors.New("flow not found :" + o.mdl.FlowId)
	}

	return flow, nil
}

func (o *FlowInstance) Execute(param map[string]interface{}) error {

	o.mdl = repo.GetInstanceRepo().GetById(o.Id)

	if o.mdl == nil {
		return errors.New("instance not found : " + o.Id)
	}

	o.Status = o.mdl.Status

	if !(o.Status == 0 || o.Status == 1) {
		return errors.New("instance is stopped")
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

	o.Status = o.mdl.Status

	if !(o.Status == 0 || o.Status == 1) {
		return errors.New("instance is stopped")
	}

	o.SetStatus(1)

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

func (o *FlowInstance) End() {

	o.Status = 99
	repo.GetInstanceRepo().SetStepStatus(o.Id, o.StepId, 99)
}

func (o *FlowInstance) Stop(reason string) {

	o.Status = -1
	repo.GetInstanceRepo().SetStatusReason(o.Id, -1, reason)
}

func (o *FlowInstance) SetStepId(stepId string) {

	o.StepId = stepId
	repo.GetInstanceRepo().SetStepId(o.Id, stepId)
}

func (o *FlowInstance) SetStatus(status int) {

	o.Status = status
	repo.GetInstanceRepo().SetStatus(o.Id, status)
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

func GetInstace(instanceId string) (*FlowInstance, error) {

	mdl := repo.GetInstanceRepo().GetById(instanceId)

	if mdl == nil {

		return nil, errors.New("instance not found")
	}

	flow, err := GetFlow(mdl.FlowId)

	if err != nil {

		return nil, err
	}

	step := flow.getStep(mdl.StepId)

	if step == nil {

		return nil, errors.New("step not found")
	}

	return &FlowInstance{
		Id:       instanceId,
		Status:   mdl.Status,
		FlowId:   flow.Id,
		FlowName: flow.Name,
		StepId:   step.Id,
		StepName: step.Name,
	}, nil
}
