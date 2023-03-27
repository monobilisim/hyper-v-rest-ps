package wmi

import (
	"os/exec"
)

func execPS(ps string) ([]byte, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", ps)
	return cmd.Output()
}
