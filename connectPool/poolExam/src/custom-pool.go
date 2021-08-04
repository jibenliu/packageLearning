package src

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var nowFunc = time.Now

type Conn struct {
	maxConn       int                     //最大可连接数
	maxIdle       int                     //最大可用连接数
	freeConn      int                     //线程池空闲连接数
	connPool      []int                   //连接池
	openCount     int                     //已经打开的连接数
	waitConn      map[int]chan Permission //排队等待的连接队列
	waitCount     int                     //等待个数
	lock          sync.Mutex              //锁
	nextConnIndex NextConnIndex           // 下一个连接的ID标识（用于区分每个ID）
	freeConns     map[int]Permission      //连接池
}

type Permission struct {
	NextConnIndex               // 对应Conn中的NextConnIndex
	Content       string        // 通行证的具体内容，比如"PASSED"表示成功获取
	CreatedAt     time.Time     // 创建时间，即连接的创建时间
	MaxLifeTime   time.Duration // 连接的存活时间，本次没有用到这个属性，保留
}

type NextConnIndex struct {
	Index int
}

type Config struct {
	MaxConn int
	MaxIdle int
}

func Prepare(ctx context.Context, config *Config) (conn *Conn) {
	// go func() {
	//for {
	//conn.expiredCh = make(chan string, len(conn.freeConns))
	//for _, value := range conn.freeConns {
	//	if value.CreatedAt.Add(value.MaxLifeTime).Before(nowFunc()) {
	//		conn.expiredCh <- "CLOSE"
	//	}
	//}
	// }()

	return &Conn{
		maxConn:   config.MaxConn,
		maxIdle:   config.MaxIdle,
		openCount: 0,
		connPool:  []int{},
		waitConn:  make(map[int]chan Permission),
		waitCount: 0,
		freeConns: make(map[int]Permission),
	}
}

func (conn *Conn) New(ctx context.Context) (permission Permission, err error) {
	/**
	1、如果当前连接池已满，即len(freeConns)=0
	2、判定openConn是否大于maxConn，如果大于，则丢弃获取加入队列进行等待
	3、如果小于，则考虑创建新连接
	*/
	conn.lock.Lock()

	select {
	default:
	case <-ctx.Done(): // context取消或超时，则退出
		conn.lock.Unlock()
		return Permission{}, errors.New("new conn failed,context cancelled")
	}
	// 连接池不为空，从连接池获取连接
	if len(conn.freeConns) > 0 {
		var (
			popPermission Permission
			popReqKey     int
		)

		//获取其中一个链接
		for popReqKey, popPermission = range conn.freeConns {
			break
		}
		//从连接池删除
		delete(conn.freeConns, popReqKey)
		fmt.Println("log", "use free conn!!!", "openCount:", conn.openCount, " freeConns: ", conn.freeConns)
		conn.lock.Unlock()
		return popPermission, nil
	}

	if conn.openCount >= conn.maxConn { // 当前连接数大于上限，则加入等待队列
		nextConnIndex := getNextConnIndex(conn)
		req := make(chan Permission, 1)
		conn.waitConn[nextConnIndex] = req
		conn.lock.Unlock()

		select {
		// 如果在等待指定超时时间后，仍然无法获取释放连接，则放弃获取连接，这里如果不在超时时间后退出会一直阻塞
		case <-time.After(time.Second * time.Duration(3)):
			fmt.Println("超时，通知主线程退出")
			break
		case ret, ok := <-req: // 有放回的连接, 直接拿来用
			if !ok {
				return Permission{}, errors.New("new conn failed, no available conn release")
			}
			fmt.Println("log", "received released conn!!!!!", "openCount: ", conn.openCount, " freeConns: ", conn.freeConns)
			return ret, nil
		}
		return Permission{}, errors.New("new conn failed")
	}
	//新建连接
	conn.openCount++
	conn.lock.Unlock()
	nextConnIndex := getNextConnIndex(conn)
	permission = Permission{NextConnIndex: NextConnIndex{nextConnIndex}, Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
	fmt.Println("log", "create conn!!!!!", "openCount: ", conn.openCount, " freeConns: ", conn.freeConns)
	return permission, nil
}
func getNextConnIndex(conn *Conn) int {
	currentIndex := conn.nextConnIndex.Index
	conn.nextConnIndex.Index = currentIndex + 1
	return conn.nextConnIndex.Index
}

func (conn *Conn) Release(ctx context.Context) (result bool, err error) {
	conn.lock.Lock()
	// 如果等待队列有等待任务，则通知正在阻塞等待获取连接的进程（即New方法中"<-req"逻辑）
	// 这里没有做指定连接的释放，只是保证释放的连接会被利用起来
	if len(conn.waitConn) > 0 {
		var req chan Permission
		var reqKey int
		for reqKey, req = range conn.waitConn {
			break
		}
		//假定释放的连接就是下面新建的连接
		permission := Permission{NextConnIndex: NextConnIndex{reqKey},
			Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
		req <-permission
		conn.waitCount--
		delete(conn.waitConn, reqKey)
		conn.lock.Unlock()
	} else {
		if conn.openCount > 0 {
			conn.openCount--
			if len(conn.freeConns) < conn.maxIdle { // 确保连接池大小不会超过maxIdle
				nextConnIndex := getNextConnIndex(conn)
				permission := Permission{NextConnIndex: NextConnIndex{nextConnIndex},
					Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
				conn.freeConns[nextConnIndex] = permission
			}
		}
		conn.lock.Unlock()
	}
	return
}
