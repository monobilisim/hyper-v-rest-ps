package utilities

import (
	"container/list"
	"time"

	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

var shellQueue *list.List
var selectedShell interface{}

func InitPwsh() {
	shellQueue = list.New()
	addShellInstance()
}

func addShellInstance() {
	shell, err := powershell.New(&backend.Local{})
	if err != nil {
		panic(err)
	}
	shellQueue.PushBack(shell)
	selectedShell = shellQueue.Front().Value
}

func CommandLine(command string) ([]byte, error) {
	if shellQueue == nil || shellQueue.Len() < 3 {
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

		if shellQueue.Len() > 0 {
			shellQueue.MoveToBack(shellQueue.Front())
			selectedShell = shellQueue.Front().Value
		}

		result <- []byte(output)
	}()

	select {
	case <-time.After(2 * time.Second):
		shell, ok := selectedShell.(powershell.Shell)
		if ok {
			go shell.Exit()
		}

		if shellQueue != nil && shellQueue.Front() != nil {
			shellQueue.Remove(shellQueue.Front())
		}

		if shellQueue.Len() > 0 {
			selectedShell = shellQueue.Front().Value
		} else {
			go addShellInstance()
		}

		return CommandLine(command)
	case err := <-errChan:
		return nil, err
	case res := <-result:
		return res, nil
	}
}
