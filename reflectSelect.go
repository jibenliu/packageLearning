package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Conn struct {
	ID string
	// 此处可以设计添加一些其他属性，例如数据库，websocket和其他配置
}

type Pool struct {
	idle chan *Conn
}

type Dispatcher struct {
	// 分发所有的连接
	pools []*Pool
}

func (s *Dispatcher) Select() (*Conn, error) {
	// 将所有的场景转化为reflect.Select属性
	cases := make([]reflect.SelectCase, len(s.pools))
	for i, pool := range s.pools {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(pool.idle),
		}
	}
	log.Print("等待所有空闲连接调度...")
	_, recv, ok := reflect.Select(cases)
	if !ok {
		return nil, errors.New("从池中获取case失败...")
	}
	conn, ok := recv.Interface().(*Conn)
	if !ok {
		return nil, errors.New("获取连接类型错误")
	}

	return conn, nil
}

func main() {
	// 初始化空池
	pools := make([]*Pool, 0)
	for i := 0; i < 4; i++ {
		p := new(Pool)
		p.idle = make(chan *Conn)
		pools = append(pools, p)
	}

	d := Dispatcher{pools: pools}

	// 通知当前连接闲置
	go func() {
		for i, pool := range pools {
			c := &Conn{ID: fmt.Sprintf("%d", i)}
			pool.idle <- c
		}
	}()

	// 等待并且从池中获取空闲连接
	for i := 0; i < 4; i++ {
		selected, err := d.Select()
		if err != nil {
			fmt.Printf("err: %#v", err)
		}
		fmt.Printf("当前连接ID: %s\n", selected.ID)
	}
}
