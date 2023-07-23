package utilities

import (
	"container/list"
	"fmt"

	"github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

type session struct {
	shell *powershell.Shell
	busy  bool
}

var (
	shellQueue *list.List
	taskQueue  chan struct{}
)

const maxQueueSize = 5

func Init() {
	shellQueue = list.New()
	taskQueue = make(chan struct{}, maxQueueSize)
}

func addSession() {
	if shellQueue.Len() < maxQueueSize {
		shell, err := powershell.New(&backend.Local{})
		if err != nil {
			panic(err)
		}
		newSession := &session{shell: &shell, busy: false}
		shellQueue.PushBack(newSession)
		// If there are sessions in the queue and taskQueue is not full, release a slot for the next task
		if taskQueue != nil && len(taskQueue) < cap(taskQueue) {
			taskQueue <- struct{}{}
			return
		}
		fmt.Println("No task queue or task queue is full")
	}
}

func rotateQueue() {
	if shellQueue.Len() > 0 {
		e := shellQueue.Front()
		shellQueue.MoveToBack(e)
	}
}

func CommandLine(command string) ([]byte, error) {
	for shellQueue.Len() < maxQueueSize {
		addSession()
	}

	if taskQueue == nil {
		return nil, fmt.Errorf("task queue is not initialized")
	}

	errChan := make(chan error)
	result := make(chan []byte)

	<-taskQueue

	for {
		if shellQueue.Len() == 0 {
			fmt.Println("No session available, waiting in the queue...")
			break
		}

		e := shellQueue.Front()
		s := e.Value.(*session)
		if s.busy {
			rotateQueue()
			continue
		}

		go func(sessionPtr *session) {
			sessionPtr.busy = true
			output, _, err := (*sessionPtr.shell).Execute(command)
			if err != nil {
				errChan <- err
				return
			}
			result <- []byte(output)
			sessionPtr.busy = false
			taskQueue <- struct{}{} // Release the taskQueue slot for the next task
		}(s)

		select {
		case err := <-errChan:
			return nil, err
		case output := <-result:
			return output, nil
		}
	}
	return nil, fmt.Errorf("no session available")
}
