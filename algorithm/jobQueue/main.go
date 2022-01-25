package jobQueue

import (
	"container/list"
	"fmt"
)

func single() {
	queue := &JobQueue{
		jobList:    list.New(),
		noticeChan: make(chan struct{}, 10),
	}

	// init worker
	workerManager := NewWorkerManager(queue)

	go workerManager.StartWork()

	// build job
	job := &SquareJob{
		BaseJob: &BaseJob{DoneChan: make(chan struct{}, 1)},
		x:       5,
	}

	// 压入队尾
	queue.PushJob(job)

	// wait for job done.
	job.WaitDone()
	fmt.Println("The End.")
}
