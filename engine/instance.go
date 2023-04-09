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

	return flow.Execute(o)
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
