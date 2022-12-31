//go:build windows

package wmi

import (
	"github.com/StackExchange/wmi"
	"time"
)

type MSVM_ComputerSystem struct {
	InstanceID                               string
	Caption                                  string
	Description                              string
	ElementName                              string
	InstallDate                              time.Time
	OperationalStatus                        []int16
	StatusDescriptions                       []string
	Status                                   string
	HealthState                              int16
	CommunicationStatus                      int16
	DetailedStatus                           int16
	OperatingStatus                          int16
	PrimaryStatus                            int16
	EnabledState                             int16
	OtherEnabledState                        string
	RequestedState                           int16
	EnabledDefault                           int16
	TimeOfLastStateChange                    time.Time
	AvailableRequestedStates                 []int16
	TransitioningToState                     int16
	CreationClassName                        string
	Name                                     string
	PrimaryOwnerName                         string
	PrimaryOwnerContact                      string
	Roles                                    []string
	NameFormat                               string
	OtherIdentifyingInfo                     []string
	IdentifyingDescriptions                  []string
	Dedicated                                []int16
	OtherDedicatedDescriptions               []string
	ResetCapability                          int16
	PowerManagementCapabilities              []int16
	OnTimeInMilliseconds                     int64
	ProcessID                                int32
	TimeOfLastConfigurationChange            time.Time
	NumberOfNumaNodes                        int16
	ReplicationState                         int16
	ReplicationHealth                        int16
	ReplicationMode                          int16
	FailedOverReplicationType                int16
	LastReplicationType                      int16
	LastApplicationConsistentReplicationTime time.Time
	LastReplicationTime                      time.Time
	LastSuccessfulBackupTime                 time.Time
	EnhancedSessionModeState                 int16
}

func VM(name string) ([]MSVM_ComputerSystem, error) {
	var v []MSVM_ComputerSystem
	q := wmi.CreateQuery(&v, "WHERE Caption='Virtual Machine' AND Name='"+name+"'")
	err := wmi.QueryNamespace(q, &v, `root\virtualization\v2`)
	if err != nil {
		return v, err
	}

	return v, err
}

func VMs() ([]MSVM_ComputerSystem, error) {
	var v []MSVM_ComputerSystem
	q := wmi.CreateQuery(&v, "WHERE Caption='Virtual Machine'")
	err := wmi.QueryNamespace(q, &v, `root\virtualization\v2`)
	if err != nil {
		return nil, err
	}

	return v, err
}
