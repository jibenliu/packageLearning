package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var ErrPoolClosed = errors.New("资源池已被关闭")

func NewPool(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size 的值太小了。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true

	close(p.res)

	for r := range p.res {
		_ = r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		_ = r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		_ = r.Close()
	}
}

/**--------------------------分割线-------------------**/
type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("关闭链接", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}

func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("第%d个查询，使用的是ID为%d的数据库链接", query, conn.(*dbConnection).ID)
}

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := NewPool(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}

	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}
