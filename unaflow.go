package unaflow

import (
	"github.com/adlindo/unaflow/dto"
	"github.com/adlindo/unaflow/engine"

	_ "github.com/adlindo/unaflow/components"
)

func ListFlow(filter string, pageNo, pageSize int) ([]dto.Flow, int64) {

	return engine.ListFlow(filter, pageNo, pageSize)
}

func GetFlow(idOrCode string) (*engine.Flow, error) {

	return engine.GetFlow(idOrCode)
}

func GetInstace(instanceId string) *engine.FlowInstance {

	return engine.GetInstace(instanceId)
}
