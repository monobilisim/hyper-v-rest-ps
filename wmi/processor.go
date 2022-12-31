//go:build windows

package wmi

import (
	"github.com/StackExchange/wmi"
	"time"
)

type MSVM_Processor struct {
	InstanceID                  string
	Caption                     string
	Description                 string
	ElementName                 string
	InstallDate                 time.Time
	Name                        string
	OperationalStatus           []int16
	StatusDescriptions          []string
	Status                      string
	HealthState                 int16
	CommunicationStatus         int16
	DetailedStatus              int16
	OperatingStatus             int16
	PrimaryStatus               int16
	EnabledState                int16
	OtherEnabledState           string
	RequestedState              int16
	EnabledDefault              int16
	TimeOfLastStateChange       time.Time
	AvailableRequestedStates    []int16
	TransitioningToState        int16
	SystemCreationClassName     string
	SystemName                  string
	CreationClassName           string
	DeviceID                    string
	PowerManagementSupported    bool
	PowerManagementCapabilities []int16
	Availability                int16
	StatusInfo                  int16
	LastErrorCode               int32
	ErrorDescription            string
	ErrorCleared                bool
	OtherIdentifyingInfo        []string
	PowerOnHours                int64
	TotalPowerOnHours           int64
	IdentifyingDescriptions     []string
	AdditionalAvailability      []int16
	MaxQuiesceTime              int64
	Role                        string
	Family                      int16
	OtherFamilyDescription      string
	UpgradeMethod               int16
	MaxClockSpeed               int32
	CurrentClockSpeed           int32
	DataWidth                   int16
	AddressWidth                int16
	LoadPercentage              int16
	Stepping                    string
	UniqueID                    string
	CPUStatus                   int16
	ExternalBusClockSpeed       int32
	LoadPercentageHistory       []int16
}

func Processor(vmName string) ([]MSVM_Processor, error) {
	var p []MSVM_Processor
	q := wmi.CreateQuery(&p, "WHERE SystemName='"+vmName+"'")
	err := wmi.QueryNamespace(q, &p, `root\virtualization\v2`)
	if err != nil {
		return nil, err
	}
	return p, err
}
