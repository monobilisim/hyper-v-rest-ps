package wmi

import (
	"github.com/StackExchange/wmi"
)

type Msvm_MemorySettingData struct {
	Reservation int64
}

func Memory(vmName string) ([]Msvm_MemorySettingData, error) {
	var dst []Msvm_MemorySettingData
	q := "ASSOCIATORS OF {Msvm_VirtualSystemSettingData.InstanceID='Microsoft:"+vmName+"'} WHERE ResultClass = Msvm_MemorySettingData"
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}
