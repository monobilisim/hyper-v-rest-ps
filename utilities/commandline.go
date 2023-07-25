package utilities

import (
	"container/list"
	"fmt"
	"sync"
	"time"

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
	Wg         *sync.WaitGroup
)

const maxQueueSize = 5

func Init() {
	shellQueue = list.New()
	taskQueue = make(chan struct{}, maxQueueSize)
	Wg = &sync.WaitGroup{}
}

func CommandLine(command string) ([]byte, error) {
	Wg.Wait()
	for shellQueue.Len() < maxQueueSize {
		addShell()
	}

	if taskQueue == nil {
		return nil, fmt.Errorf("task queue is not initialized")
	}

	return dispatchTask(command)
}

func dispatchTask(command string) ([]byte, error) {
	errChan := make(chan error)
	result := make(chan []byte)

	<-taskQueue

	for {
		if shellQueue.Len() == 0 {
			log.Warn("No session available, waiting in the queue...")
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
		case <-time.After(1080 * time.Second):
			s.busy = false
			(*s.shell).Exit()
			go RefreshShellQueue()
			return nil, fmt.Errorf("timeout")
		case err := <-errChan:
			return nil, err
		case output := <-result:
			return output, nil
		}
	}
	return nil, fmt.Errorf("no session available")
}

func addShell() {
	Wg.Wait()
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
		log.Warn("No task queue or task queue is full")
	}
}

func rotateQueue() {
	if shellQueue.Len() > 0 {
		e := shellQueue.Front()
		shellQueue.MoveToBack(e)
	}
}

func RefreshShellQueue() {
	Wg.Add(1)
	defer Wg.Done()
	for shellQueue.Len() > 1 {
		e := shellQueue.Front()
		s := e.Value.(*session)
		if s.busy {
			rotateQueue()
			continue
		}
		(*s.shell).Exit()
		shellQueue.Remove(e)
	}
}
