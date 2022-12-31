//go:build windows

package wmi

import (
	"github.com/StackExchange/wmi"
	"time"
)

type MSVM_DiskDrive struct {
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
	Capabilities                []int16
	CapabilityDescriptions      []string
	ErrorMethodology            string
	CompressionMethod           string
	NumberOfMediaSupported      int32
	MaxMediaSize                int64
	DefaultBlockSize            int64
	MaxBlockSize                int64
	MinBlockSize                int64
	NeedsCleaning               bool
	MediaIsLocked               bool
	Security                    int16
	LastCleaned                 time.Time
	MaxAccessTime               int64
	UncompressedDataRate        int32
	LoadTime                    int64
	UnloadTime                  int64
	MountCount                  int64
	TimeOfLastMount             time.Time
	TotalMountTime              int64
	UnitsDescription            string
	MaxUnitsBeforeCleaning      uint64
	UnitsUsed                   int64
	DriveNumber                 int32
}

func Storage(vmName string) ([]MSVM_DiskDrive, error) {
	var s []MSVM_DiskDrive
	q := wmi.CreateQuery(&s, "WHERE SystemName='"+vmName+"'")
	err := wmi.QueryNamespace(q, &s, `root\virtualization\v2`)
	if err != nil {
		return nil, err
	}
	return s, err
}
