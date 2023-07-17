package hyperv

import "time"

type Msvm_ComputerSystem struct {
	ElementName string
	InstallDate time.Time
	Name        string
	ProcessID   int32
}

type VHDData struct {
	ComputerName            string      `json:"ComputerName"`
	Path                    string      `json:"Path"`
	VhdFormat               int64       `json:"VhdFormat"`
	VhdType                 int64       `json:"VhdType"`
	FileSize                int64       `json:"FileSize"`
	Size                    int64       `json:"Size"`
	MinimumSize             interface{} `json:"MinimumSize"`
	LogicalSectorSize       int64       `json:"LogicalSectorSize"`
	PhysicalSectorSize      int64       `json:"PhysicalSectorSize"`
	BlockSize               int64       `json:"BlockSize"`
	ParentPath              string      `json:"ParentPath"`
	DiskIdentifier          string      `json:"DiskIdentifier"`
	FragmentationPercentage int64       `json:"FragmentationPercentage"`
	Alignment               int64       `json:"Alignment"`
	Attached                bool        `json:"Attached"`
	DiskNumber              interface{} `json:"DiskNumber"`
	IsPMEMCompatible        bool        `json:"IsPMEMCompatible"`
	AddressAbstractionType  int64       `json:"AddressAbstractionType"`
	Number                  interface{} `json:"Number"`
}

type MemoryData struct {
	ResourcePoolName        string      `json:"ResourcePoolName"`
	Buffer                  int64       `json:"Buffer"`
	DynamicMemoryEnabled    bool        `json:"DynamicMemoryEnabled"`
	Maximum                 int64       `json:"Maximum"`
	MaximumPerNUMANode      int64       `json:"MaximumPerNumaNode"`
	Minimum                 int64       `json:"Minimum"`
	Priority                int64       `json:"Priority"`
	Startup                 int64       `json:"Startup"`
	HugePagesEnabled        bool        `json:"HugePagesEnabled"`
	MemoryEncryptionPolicy  int64       `json:"MemoryEncryptionPolicy"`
	MemoryEncryptionEnabled interface{} `json:"MemoryEncryptionEnabled"`
	Name                    string      `json:"Name"`
	ID                      string      `json:"Id"`
	VMID                    string      `json:"VMId"`
	VMName                  string      `json:"VMName"`
	VMSnapshotID            string      `json:"VMSnapshotId"`
	VMSnapshotName          string      `json:"VMSnapshotName"`
	CIMSession              CIMSession  `json:"CimSession"`
	ComputerName            string      `json:"ComputerName"`
	IsDeleted               bool        `json:"IsDeleted"`
	VMCheckpointID          string      `json:"VMCheckpointId"`
	VMCheckpointName        string      `json:"VMCheckpointName"`
}

type CIMSession struct {
	ComputerName interface{} `json:"ComputerName"`
	InstanceID   string      `json:"InstanceId"`
}

type ProcessorData struct {
	ResourcePoolName                             string        `json:"ResourcePoolName"`
	Count                                        int64         `json:"Count"`
	CompatibilityForMigrationEnabled             bool          `json:"CompatibilityForMigrationEnabled"`
	CompatibilityForMigrationMode                int64         `json:"CompatibilityForMigrationMode"`
	CompatibilityForOlderOperatingSystemsEnabled bool          `json:"CompatibilityForOlderOperatingSystemsEnabled"`
	HwThreadCountPerCore                         int64         `json:"HwThreadCountPerCore"`
	ExposeVirtualizationExtensions               bool          `json:"ExposeVirtualizationExtensions"`
	EnablePerfmonPmu                             bool          `json:"EnablePerfmonPmu"`
	EnablePerfmonLbr                             bool          `json:"EnablePerfmonLbr"`
	EnablePerfmonPebs                            bool          `json:"EnablePerfmonPebs"`
	EnablePerfmonIpt                             bool          `json:"EnablePerfmonIpt"`
	EnableLegacyAPICMode                         bool          `json:"EnableLegacyApicMode"`
	APICMode                                     int64         `json:"ApicMode"`
	AllowACountMCount                            bool          `json:"AllowACountMCount"`
	CPUBrandString                               string        `json:"CpuBrandString"`
	PerfCPUFreqCapMhz                            int64         `json:"PerfCpuFreqCapMhz"`
	Maximum                                      int64         `json:"Maximum"`
	Reserve                                      int64         `json:"Reserve"`
	RelativeWeight                               int64         `json:"RelativeWeight"`
	MaximumCountPerNUMANode                      int64         `json:"MaximumCountPerNumaNode"`
	MaximumCountPerNUMASocket                    int64         `json:"MaximumCountPerNumaSocket"`
	EnableHostResourceProtection                 bool          `json:"EnableHostResourceProtection"`
	OperationalStatus                            []interface{} `json:"OperationalStatus"`
	StatusDescription                            []interface{} `json:"StatusDescription"`
	Name                                         string        `json:"Name"`
	ID                                           string        `json:"Id"`
	VMID                                         string        `json:"VMId"`
	VMName                                       string        `json:"VMName"`
	VMSnapshotID                                 string        `json:"VMSnapshotId"`
	VMSnapshotName                               string        `json:"VMSnapshotName"`
	CIMSession                                   CIMSession    `json:"CimSession"`
	ComputerName                                 string        `json:"ComputerName"`
	IsDeleted                                    bool          `json:"IsDeleted"`
	VMCheckpointID                               string        `json:"VMCheckpointId"`
	VMCheckpointName                             string        `json:"VMCheckpointName"`
}
