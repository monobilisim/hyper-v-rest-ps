package wmi

import (
	"github.com/StackExchange/wmi"
	"time"
)

type MSVM_ComputerSystem struct {
	ElementName                              string
	InstallDate                              time.Time
	Name                                     string
	ProcessID                                int32
}

func VMs() ([]MSVM_ComputerSystem, error) {
	var dst []MSVM_ComputerSystem
	q := wmi.CreateQuery(&dst, "WHERE Caption='Virtual Machine'")
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}
