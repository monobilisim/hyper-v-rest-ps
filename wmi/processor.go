package wmi

import (
	"github.com/StackExchange/wmi"
)

type Msvm_SummaryInformation struct {
	NumberOfProcessors   int16
	GuestOperatingSystem string
}

func Processor(vmName string) ([]Msvm_SummaryInformation, error) {
	var dst []Msvm_SummaryInformation
	q := wmi.CreateQuery(&dst, "WHERE Name = '" + vmName + "'")
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}
