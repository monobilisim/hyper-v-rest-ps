package wmi

import (
	"fmt"
	"log"
)

// Vhd retrieves information about the virtual hard disk associated with the virtual machine identified by the given VMId.
//
// Parameters:
// VMId - The unique identifier of the virtual machine.
//
// Returns:
// A JSON-encoded byte slice containing information about the virtual hard disk.
// Any errors encountered during execution are returned as an error.
func Vhd(VMId string) ([]byte, error) {
	ps := `Get-VHD -Id ` + VMId + ` | ConvertTo-Json`
	output, err := execPS(ps)
	if err != nil {
		log.Printf("Failed to retrieve VHD: %v", err)
		return nil, fmt.Errorf("failed to retrieve VHD: %v", err)
	}
	return output, nil
}
