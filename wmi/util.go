package wmi

import (
	"os/exec"
)

// execPS executes the given PowerShell command and returns its output as a byte slice.
//
// Parameters:
// ps - The PowerShell command to execute.
//
// Returns:
// The output of the PowerShell command as a byte slice.
// Any errors encountered during execution are returned as an error.
func execPS(ps string) ([]byte, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", ps)
	return cmd.Output()
}
