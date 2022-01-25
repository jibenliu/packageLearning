package jobQueue

import (
	"fmt"
)

type WorkerManager struct {
	queue     *JobQueue
	closeChan chan struct{}
}

func (m *WorkerManager) StartWork() error {
	fmt.Println("Start to work...")
	for {
		select {
		case <-m.closeChan:
			return nil
		case <-m.queue.noticeChan:
			job := m.queue.PopJob()
			m.ConsumeJob(job)
		}
	}
}

func (m *WorkerManager) ConsumeJob(job Job) {
	defer func() {
		job.Done()
	}()

	job.Execute()
}

func NewWorkerManager(queue *JobQueue) WorkerManager {
	return WorkerManager{
		queue: queue,
	}
}
