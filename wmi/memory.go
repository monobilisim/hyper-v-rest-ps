//go:build windows

package wmi

import (
	"github.com/StackExchange/wmi"
	"time"
)

type MSVM_Memory struct {
	Caption                       string
	Description                   string
	ElementName                   string
	InstallDate                   time.Time
	OperationalStatus             []int16
	Status                        string
	HealthState                   int16
	EnabledState                  int16
	OtherEnabledState             string
	RequestedState                int16
	TimeOfLastStateChange         time.Time
	SystemCreationClassName       string
	SystemName                    string
	PowerManagementSupported      bool
	PowerManagementCapabilities   []int16
	Availability                  int16
	StatusInfo                    int16
	LastErrorCode                 int32
	ErrorDescription              string
	ErrorCleared                  bool
	PowerOnHours                  int64
	TotalPowerOnHours             int64
	IdentifyingDescriptions       []string
	AdditionalAvailability        []int16
	MaxQuiesceTime                int64
	DataOrganization              int16
	Purpose                       string
	Access                        int16
	NumberOfBlocks                int64
	IsBasedOnUnderlyingRedundancy bool
	SequentialAccess              bool
	ExtentStatus                  []int16
	NoSinglePointOfFailure        bool
	DataRedundancy                int16
	PackageRedundancy             int16
	DeltaReservation              int8
	Primordial                    bool
	Name                          string
	NameFormat                    int16
	NameNamespace                 int16
	ErrorMethodology              string
	StartingAddress               int64
	ErrorInfo                     int16
	OtherErrorDescription         string
	CorrectableError              bool
	ErrorTime                     time.Time
	ErrorAccess                   int16
	ErrorTransferSize             int32
	ErrorData                     []int8
	ErrorDataOrder                int16
	ErrorAddress                  int64
	SystemLevelAddress            bool
	ErrorResolution               int64
	AdditionalErrorData           []int8
	StatusDescriptions            []string
	EnabledDefault                int16
	CreationClassName             string
	DeviceID                      string
	OtherIdentifyingInfo          []string
	BlockSize                     int64
	ConsumableBlocks              int64
	OtherNameNamespace            string
	OtherNameFormat               string
	EndingAddress                 int64
}

func Memory(vmName string) ([]MSVM_Memory, error) {
	var m []MSVM_Memory
	q := wmi.CreateQuery(&m, "WHERE SystemName='"+vmName+"'")
	err := wmi.QueryNamespace(q, &m, `root\virtualization\v2`)
	if err != nil {
		return nil, err
	}
	return m, err
}
