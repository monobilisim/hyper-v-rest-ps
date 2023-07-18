package utilities

import (
	"time"

	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

var shell powershell.Shell
var backupShell powershell.Shell
var selectedShell int

func InitPwsh() {
	var err error
	shell, err = powershell.New(&backend.Local{})
	if err != nil {
		panic(err)
	}
}

func CommandLine(command string) ([]byte, error) {
	result := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		var output string
		var err error
		if selectedShell == 0 {
			output, _, err = shell.Execute(command)
		} else {
			output, _, err = backupShell.Execute(command)
		}

		if err != nil {
			errChan <- err
			return
		}
		result <- []byte(output)
	}()

	select {
	case <-time.After(2 * time.Second):
		if selectedShell == 0 {
			shell.Exit()
			InitPwsh()
			selectedShell = 1
		} else {
			backupShell.Exit()
			backupShell, _ = powershell.New(&backend.Local{})
		}
		return CommandLine(command)
	case err := <-errChan:
		return nil, err
	case res := <-result:
		return res, nil
	}
}
