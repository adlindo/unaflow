package unaflow

import (
	"strings"
	"sync"
	"time"

	"github.com/adlindo/gocom"
	"github.com/adlindo/unaflow/repo"
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

	gocom.GET(o.basePath+"/flow", o.listFlow)
}

func (o *UnaFlowCtrl) test(ctx gocom.Context) error {

	return ctx.SendString("OK " + time.Now().Format("2006-01-02 15:04:05"))
}

func (o *UnaFlowCtrl) listFlow(ctx gocom.Context) error {

	ret := ListFlow()

	return ctx.SendResult(ret)
}
