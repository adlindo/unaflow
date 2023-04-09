package main

import (
	"fmt"

	"github.com/adlindo/gocom"
	"github.com/adlindo/unaflow"
)

func main() {

	gocom.AddCtrl(unaflow.GetUnaFlowCtrl())

	flow, err := unaflow.GetFlow("a")

	if err != nil {

		fmt.Println("Error get from ", err)
	}

	fmt.Println("sebelummmmmmmmmmmmmmmmmmmmm")
	fmt.Println("====>", flow)

	gocom.Start()
}
