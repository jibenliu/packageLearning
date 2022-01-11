package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

//Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//从操作系统发送信号
	interrupt chan os.Signal
	//报告处理任务已完成
	complete chan error
	//报告处理任务已经超时
	timeout <-chan time.Time
	//持有一组以索引为顺序的依次执行的以 int 类型 id 为参数的函数
	tasks []func(id int)
}

//统一错误处理

// ErrTimeOut 超时错误信息，会在任务执行超时时返回
var ErrTimeOut = errors.New("received timeout")

// ErrInterrupt 中断错误信号，会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

//New 函数返回一个新的准备使用的 Runner，d：自定义分配的时间
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		//会在另一线程经过时间段 d 后向返回值发送当时的时间。
		timeout: time.After(d),
	}
}

//Add 将一个任务加入到 Runner 中
func (r *Runner) Add(tasks ...func(id int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start 开始执行所有任务，并监控通道事件
func (r *Runner) Start() error {
	//监控所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	//使用不同的 goroutine 执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	//使用 select 语句来监控 goroutine 的通信
	select {
	//等待任务完成
	case err := <-r.complete:
		return err
	//任务超时
	case <-r.timeout:
		return ErrTimeOut
	}
}

//执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//执行已注册的任务
		task(id)
	}
	return nil
}

//检测是否收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//当中断事件被触发时
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
		//继续执行
	default:
		return false
	}
}

const timeout = 1 * time.Second

func main() {
	log.Println("Starting work.")

	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt.")
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor -> Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
