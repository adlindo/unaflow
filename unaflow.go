package unaflow

import (
	"github.com/adlindo/unaflow/dto"
	"github.com/adlindo/unaflow/engine"

	_ "github.com/adlindo/unaflow/components"
)

func ListFlow() []dto.Flow {

	ret := []dto.Flow{}

	return ret
}

func GetFlow(flowId string) (*engine.Flow, error) {

	return engine.GetFlow(flowId)
}

func GetInstace(instanceId string) *engine.FlowInstance {

	return engine.GetInstace(instanceId)
}
