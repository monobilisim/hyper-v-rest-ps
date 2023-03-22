package wmi

import (
	"fmt"
	"os/exec"
)

func execPS(ps string) ([]byte, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", ps)
	fmt.Println("Executing: ", cmd.Args)

	return cmd.Output()
}
