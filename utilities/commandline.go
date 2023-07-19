package utilities

import (
	"container/list"
	"time"

	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

var shellQueue *list.List
var selectedShell interface{}

func CommandLine(command string) ([]byte, error) {
	if shellQueue == nil || shellQueue.Len() < 4 {
		go addShellInstance()
		go addShellInstance()
	}

	errChan := make(chan error)
	result := make(chan []byte)

	go func() {
		var output string

		shell, ok := selectedShell.(powershell.Shell)
		if ok {
			var err error
			output, _, err = shell.Execute(command)
			if err != nil {
				errChan <- err
				return
			}
		}

		result <- []byte(output)
	}()
	rotateShell()
	select {
	case <-time.After(20 * time.Second):
		shell, ok := selectedShell.(powershell.Shell)
		if ok {
			shell.Exit()
		}

		if shellQueue != nil && shellQueue.Front() != nil {
			shellQueue.Remove(shellQueue.Front())
		}

		if shellQueue.Len() > 0 {
			selectedShell = shellQueue.Front().Value
		}

		rotateShell()

		return CommandLine(command)
	case err := <-errChan:
		return nil, err
	case res := <-result:
		return res, nil
	}
}

func InitPwsh() {
	shellQueue = list.New()
	addShellInstance()
}

func rotateShell() {
	if shellQueue.Len() > 1 {
		shellQueue.MoveToBack(shellQueue.Front())
		selectedShell = shellQueue.Front().Value
	}
}

func addShellInstance() {
	if shellQueue.Len() < 4 {
		shell, err := powershell.New(&backend.Local{})
		if err != nil {
			panic(err)
		}
		shellQueue.PushBack(shell)
		selectedShell = shellQueue.Front().Value
	}
}
