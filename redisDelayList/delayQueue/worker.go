package delayQueue

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// TopicWorker
type TopicWorker struct {
	TopicName   string              // topic 名称
	Handler     JobHandler          // 处理job的方法
	WorkerCount int                 // 并行任务数
	WorkerPool  *semaphore.Weighted // 通过信号量控制并发协程数
}

func NewTopicWorker(handler JobHandler, workerCount int) *TopicWorker {
	return &TopicWorker{Handler: handler, WorkerCount: workerCount, WorkerPool: semaphore.NewWeighted(int64(workerCount)), TopicName: handler.Topic()}
}

// JobHandler
type JobHandler interface {
	Topic() string                         // 返回topic名称
	Execute(context.Context, []byte) error // 处理消息
	ErrorHandle(job *Job) error            //失败处理函数
}

type PrintJob struct {
}

func (p *PrintJob) Topic() string {
	return "PrintJob"
}

func (p *PrintJob) Execute(ctx context.Context, bytes []byte) error {
	fmt.Printf("执行任务:%s\n", bytes)
	return nil
}

func (p *PrintJob) ErrorHandle(job *Job) error {
	//超过10次扔掉
	if job.TryCount > 10 {
		q := queue{}
		job.Delay = time.Now().Unix() + job.TryTimeOut //超时后一定时间再次执行
		job.TryCount += 1
		return q.DelayJob(job)
	}
	return nil
}
