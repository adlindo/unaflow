package main

import (
	"github.com/adlindo/gocom"
	"github.com/adlindo/unaflow"
)

func main() {

	gocom.AddCtrl(unaflow.GetUnaFlowCtrl())
	gocom.Start()
}
