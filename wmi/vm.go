package wmi

import (
	"fmt"
	"log"
)

// VMs retrieves information about all virtual machines managed by the hypervisor.
//
// Returns:
// A JSON-encoded byte slice containing information about all virtual machines.
// Any errors encountered during execution are returned as an error.
func VMs() ([]byte, error) {
	ps := `Get-VM | ConvertTo-Json`
	output, err := execPS(ps)
	if err != nil {
		log.Printf("Failed to retrieve VMs: %v", err)
		return nil, fmt.Errorf("failed to retrieve VMs: %v", err)
	}
	return output, nil
}

// VM retrieves information about a specific virtual machine identified by the given VMId.
//
// Parameters:
// VMId - The unique identifier of the virtual machine.
//
// Returns:
// A JSON-encoded byte slice containing information about the virtual machine.
// Any errors encountered during execution are returned as an error.
func VM(VMId string) ([]byte, error) {
	ps := `Get-VM -Id ` + VMId + ` | ConvertTo-Json`
	output, err := execPS(ps)
	if err != nil {
		log.Printf("Failed to retrieve VM: %v", err)
		return nil, fmt.Errorf("failed to retrieve VM: %v", err)
	}
	return output, nil
}
