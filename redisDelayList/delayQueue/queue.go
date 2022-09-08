package delayQueue

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

const closed = 1
const DelayJobType = "DelayJobType"
const ReadyJobType = "ReadyJobType"

var queueServer *QueueServer

func InitQueue(ctx context.Context) {
	go func() {
		queueServer = NewQueueServer(&QueueOption{CloseWaitTime: 7, WaitTime: 2})
		queueServer.Register(NewTopicWorker(&PrintJob{}, 10))
		err := queueServer.Run(ctx)
		if err != nil {
			ctx.Done()
			panic(err)
		}
	}()
}

type QueueOption struct {
	CloseWaitTime int64
	WaitTime      int64
}

type QueueServer struct {
	queueOption *QueueOption
	topicWokers []*TopicWorker
	queue       *queue
	stopCh      chan struct{}
	close       uint32
}

func NewQueueServer(queueOption *QueueOption) *QueueServer {
	return &QueueServer{queueOption: queueOption, queue: &queue{}, topicWokers: make([]*TopicWorker, 0), stopCh: make(chan struct{}, 1)}
}

func (q *QueueServer) Register(work *TopicWorker) {
	q.topicWokers = append(q.topicWokers, work)
}

// QueueClient
type QueueClient struct {
	queue *queue
}

// Run 开始处理消息直到收到退出命令
func (s *QueueServer) Run(ctx context.Context) error {
	// 一个topic 一个协程处理
	for _, topicWorker := range s.topicWokers {
		go s.processJob(ctx, topicWorker)
	}
	// 监听系统信号
	go s.watchSystemSignal(ctx)
	// 等待退出
	<-s.stopCh
	log.Infof("server will close")
	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, time.Second*time.Duration(s.queueOption.CloseWaitTime))
	defer cancelFunc()
	return s.Close(ctxTimeOut)
}

func (s *QueueServer) processJob(ctx context.Context, topic *TopicWorker) error {
	// 循环获取消息直到server退出
	for {
		// 判断server是否退出
		if atomic.LoadUint32(&s.close) == closed {
			break
		}
		// 通过信号量控制并发的协程数，正在运行的协程达到上限就等待
		if err := topic.WorkerPool.Acquire(ctx, 1); err != nil {
			return err
		}
		rdNum, dlNum, err := s.queue.GetJobNum(topic.TopicName)
		if err != nil {
			return err
		}
		log.Infof("Topic:%s ReadyJobNum: %d DelayJobNum:%d\n", topic.TopicName, rdNum, dlNum)
		job, err := s.queue.GetReadyJob(topic.TopicName)
		if err != nil && err != redis.Nil {
			topic.WorkerPool.Release(1)
			continue
		}
		if job == nil {
			topic.WorkerPool.Release(1)
			time.Sleep(time.Second * time.Duration(s.queueOption.WaitTime))
			continue
		}
		log.Infof("Topic:%s get job  %v\n", topic.TopicName, job)
		go func() {
			err := topic.Handler.Execute(ctx, job.PlayLoad)
			if err != nil {
				err = topic.Handler.ErrorHandle(job)
				if err != nil {
					log.Errorf("error handle %v %v", err, job)
				}
			}
			topic.WorkerPool.Release(1)
		}()
	}
	// 等待所有处理消息的协程执行完成
	if err := topic.WorkerPool.Acquire(ctx, int64(topic.WorkerCount)); err != nil {
		return err
	}
	log.Infof("exit processJob goruntine")
	return nil
}

func (s *QueueServer) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	atomic.StoreUint32(&s.close, closed)
	ctx.Done()
	return nil
}

func (s *QueueServer) watchSystemSignal(ctx context.Context) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGKILL)
	for {
		select {
		case <-ctx.Done():
			s.stopCh <- struct{}{}
			break
		case <-interrupt:
			log.Infof("receiver close single,server will be close")
			s.stopCh <- struct{}{}
			break
		default:
			runtime.Gosched()
		}
	}

}

type queue struct{}

func (q *queue) Push(j *Job) error {
	jMar, err := msgpack.Marshal(j)
	if err != nil {
		return err
	}
	return lpush(GetReadyKey(j.Topic), jMar)
}

func (q *queue) DelayJob(j *Job) error {

	jMar, err := msgpack.Marshal(j)
	if err != nil {
		return err
	}
	return zadd(GetDelayKey(j.Topic), jMar, int(j.Delay))
}

func GetReadyKey(topic string) string {
	return fmt.Sprintf("%s:%s", topic, ReadyJobType)
}
func GetDelayKey(topic string) string {
	return fmt.Sprintf("%s:%s", topic, DelayJobType)
}

func (q *queue) GetJobNum(topic string) (int, int, error) {
	rdNum, err := llen(GetReadyKey(topic)).Result()
	if err != nil {
		return 0, 0, err
	}
	dlNum, err := zcard(GetDelayKey(topic)).Result()
	return int(rdNum), int(dlNum), err
}

// GetReadyJob 把延迟队列的任务 写进就绪队列 拉取就绪队列任务
func (q *queue) GetReadyJob(topic string) (*Job, error) {

	err := migrateExpiredJobs(
		GetDelayKey(topic),
		GetReadyKey(topic))
	if err != nil {
		return nil, err
	}
	data, err := rpop(GetReadyKey(topic)).Result()
	if err != nil {
		return nil, nil
	}
	res := &Job{}
	err = msgpack.Unmarshal([]byte(data), res)
	return res, err
}
