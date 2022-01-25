package jobQueue

import (
	"container/list"
	"sync"
)

type JobQueue struct {
	mtx        sync.Mutex
	jobList    *list.List
	noticeChan chan struct{}
}

// 队列的Push
func (queue *JobQueue) PushJob(job Job) {
	// 将job添加到队列中
	queue.jobList.PushBack(job)
	queue.noticeChan <- struct{}{}
}

// 队列的Pop
func (queue *JobQueue) PopJob() Job {
	queue.mtx.Lock()
	defer queue.mtx.Unlock()

	// 没元素则返回nil
	if queue.jobList.Len() == 0 {
		return nil
	}

	elements := queue.jobList.Front()
	return queue.jobList.Remove(elements).(Job) // 移除并返回
}

func (queue *JobQueue) WaitJob() <-chan struct{} {
	return queue.noticeChan
}
