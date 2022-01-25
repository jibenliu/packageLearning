package jobQueue

import (
	"context"
	"fmt"
)

type Job interface {
	Execute() error
	WaitDone()
	Done()
}

type BaseJob struct {
	Err        error
	DoneChan   chan struct{}
	Ctx        context.Context
	cancelFunc context.CancelFunc
}

func (job *BaseJob) Done() {
	close(job.DoneChan)
}

func (job *BaseJob) WaitDone() {
	select {
	case <-job.DoneChan:
		return
	}
}

type SquareJob struct {
	*BaseJob
	x int
}

func (s *SquareJob) Execute() error {
	result := s.x * s.x
	fmt.Println("the result is: ", result)
	return nil
}
