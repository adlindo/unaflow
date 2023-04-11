package unaflow

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adlindo/gocom"
	"github.com/adlindo/unaflow/dto"
	"github.com/adlindo/unaflow/engine"
	"github.com/adlindo/unaflow/repo"
	"github.com/jinzhu/copier"
)

type UnaFlowCtrl struct {
	basePath string
}

var unaFlowCtrl *UnaFlowCtrl
var unaFlowCtrlOnce sync.Once

func GetUnaFlowCtrl(basePath ...string) *UnaFlowCtrl {

	unaFlowCtrlOnce.Do(func() {

		bp := ""

		if len(basePath) > 0 {
			bp = basePath[0]
		}

		if strings.HasSuffix(bp, "/") {

			bp = bp[0 : len(bp)-1]
		}

		if bp == "" {

			bp = "/unaflow"
		}

		unaFlowCtrl = &UnaFlowCtrl{
			basePath: bp,
		}

		repo.GetFlowRepo()
		repo.GetInstanceRepo()
		repo.GetInstanceDataRepo()
	})

	return unaFlowCtrl
}

func (o *UnaFlowCtrl) Init() {

	gocom.GET(o.basePath+"/test", o.test)
	gocom.GET(o.basePath+"/clear", o.clear)

	gocom.POST(o.basePath+"/flow/:idOrCode/new", o.createInstance)
	gocom.GET(o.basePath+"/flow/:idOrCode/instance", o.listInstance)
	gocom.GET(o.basePath+"/flow", o.listFlow)

	gocom.GET(o.basePath+"/instance/:id/data", o.getInstaceData)
	gocom.POST(o.basePath+"/instance/:id/data", o.setInstaceData)
	gocom.POST(o.basePath+"/instance/:id/execute", o.instanceExecute)
	gocom.POST(o.basePath+"/instance/:id/next", o.instanceNext)
	gocom.POST(o.basePath+"/instance/:id/stop", o.instanceStop)

	gocom.GET(o.basePath+"/instance/:id", o.getInstance)
}

func (o *UnaFlowCtrl) test(ctx gocom.Context) error {

	return ctx.SendString("OK " + time.Now().Format("2006-01-02 15:04:05"))
}

func (o *UnaFlowCtrl) clear(ctx gocom.Context) error {

	engine.ClearFlow()
	return ctx.SendResult(true)
}

func (o *UnaFlowCtrl) listFlow(ctx gocom.Context) error {

	pageNo, _ := strconv.Atoi(ctx.Query("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize", "10"))

	ret, total := engine.ListFlow(ctx.Query("filter"), pageNo, pageSize)

	return ctx.SendPaged(ret, pageNo, totalPage(int(total), pageSize))
}

func (o *UnaFlowCtrl) normalizeParam(param map[string]interface{}) map[string]interface{} {

	for name, val := range param {
		switch val.(type) {
		case float64:
			if val == math.Trunc(val.(float64)) {
				fmt.Println()
				param[name] = int64(val.(float64))
			}
		case map[string]interface{}:
			param[name] = o.normalizeParam(val.(map[string]interface{}))
		}
	}

	return param
}

func (o *UnaFlowCtrl) createInstance(ctx gocom.Context) error {

	data := dto.CreateInstaceReq{}
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	data.Params = o.normalizeParam(data.Params)

	flow, err := GetFlow(ctx.Param("idOrCode"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1002, err.Error()))
	}

	instance, err := flow.CreateInstance(data.Key, data.Params, data.Execute)

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	ret := dto.Instance{}
	ret.Data = instance.ListData()
	copier.Copy(&ret, instance)

	return ctx.SendResult(ret)
}

func totalPage(total, pageSize int) int {

	if pageSize < 1 {
		return 1
	}

	rem := total % pageSize

	if rem > 0 {
		return (total / pageSize) + 1
	}

	return total / pageSize
}

func (o *UnaFlowCtrl) listInstance(ctx gocom.Context) error {

	flow, err := engine.GetFlow(ctx.Param("idOrCode"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	pageNo, _ := strconv.Atoi(ctx.Query("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize", "10"))

	ret, total := flow.ListInstance(ctx.Query("step"), ctx.Query("filter"), pageNo, pageSize)

	return ctx.SendPaged(ret, pageNo, totalPage(total, pageSize))
}

//--------------------------------------------------------

func (o *UnaFlowCtrl) instanceNext(ctx gocom.Context) error {

	data := dto.NextReq{}
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	data.Params = o.normalizeParam(data.Params)

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	err = instance.Next(data.Params)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	ret := dto.Instance{}
	copier.Copy(&ret, instance)
	ret.Data = instance.ListData()

	return ctx.SendResult(ret)
}

func (o *UnaFlowCtrl) instanceExecute(ctx gocom.Context) error {

	data := dto.NextReq{}
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	data.Params = o.normalizeParam(data.Params)

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	err = instance.Execute(data.Params)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	ret := dto.Instance{}
	copier.Copy(&ret, instance)
	ret.Data = instance.ListData()

	return ctx.SendResult(ret)
}

func (o *UnaFlowCtrl) getInstance(ctx gocom.Context) error {

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	ret := dto.Instance{}
	copier.Copy(&ret, instance)
	ret.Data = instance.ListData()

	return ctx.SendResult(ret)
}

func (o *UnaFlowCtrl) getInstaceData(ctx gocom.Context) error {

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	ret := instance.ListData()

	return ctx.SendResult(ret)
}

func (o *UnaFlowCtrl) setInstaceData(ctx gocom.Context) error {

	data := dto.NextReq{}
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	data.Params = o.normalizeParam(data.Params)

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	instance.SetDataBulk(data.Params)

	return ctx.SendResult(true)
}

func (o *UnaFlowCtrl) instanceStop(ctx gocom.Context) error {

	data := dto.StopReq{}
	err := ctx.Bind(&data)

	if err != nil {
		return ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	instance, err := engine.GetInstace(ctx.Param("id"))

	if err != nil {
		return ctx.SendError(gocom.NewError(1003, err.Error()))
	}

	instance.Stop(data.Reason)

	return ctx.SendResult(true)
}
