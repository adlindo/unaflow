package unaflow

import (
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

	gocom.POST(o.basePath+"/flow/:idOrCode/new", o.createInstance)
	gocom.GET(o.basePath+"/flow/:idOrCode/instance", o.listInstance)
	gocom.GET(o.basePath+"/flow", o.listFlow)

	gocom.GET(o.basePath+"/instance/:id/data", o.listFlow)
	gocom.POST(o.basePath+"/instance/:id/data", o.listFlow)
	gocom.POST(o.basePath+"/instance/:id/next", o.listFlow)

	gocom.GET(o.basePath+"/instance/:id", o.listFlow)
}

func (o *UnaFlowCtrl) test(ctx gocom.Context) error {

	return ctx.SendString("OK " + time.Now().Format("2006-01-02 15:04:05"))
}

func (o *UnaFlowCtrl) listFlow(ctx gocom.Context) error {

	pageNo, _ := strconv.Atoi(ctx.Query("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize", "10"))

	ret, total := engine.ListFlow(ctx.Query("filter"), pageNo, pageSize)

	return ctx.SendPaged(ret, pageNo, int(total))
}

func (o *UnaFlowCtrl) createInstance(ctx gocom.Context) error {

	data := dto.CreateInstaceReq{}
	err := ctx.Bind(&data)

	if err != nil {
		ctx.SendError(gocom.NewError(1001, err.Error()))
	}

	flow, err := GetFlow(ctx.Param("idOrCode"))

	if err != nil {
		ctx.SendError(gocom.NewError(1002, err.Error()))
	}

	instance := flow.CreateInstance(data.Params, data.Execute)

	ret := dto.FlowInstance{}
	copier.Copy(&ret, instance)

	return ctx.SendResult(ret)
}

func (o *UnaFlowCtrl) listInstance(ctx gocom.Context) error {

	ret := ""

	step := ctx.Query("step")

	return ctx.SendResult(ret)
}
