package concurrency

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Person struct {
	name string
	age  int
}

// 全局变量（简单处理）
var p atomic.Value

func updateNotOk(name string, age int) { //并发场景下会导致不能同步更改
	lp := &Person{}
	// 更新第一个字段
	lp.name = name
	// 加点随机性
	time.Sleep(time.Millisecond * 200)
	// 更新第二个字段
	lp.age = age
}

func update(name string, age int) {
	lp := &Person{}
	// 更新第一个字段
	lp.name = name
	// 加点随机性
	time.Sleep(time.Millisecond * 200)
	// 更新第二个字段
	lp.age = age
	// 原子设置到全局变量
	p.Store(lp)
}

func read() {
	// 结果是啥？你能猜到吗？
	_p := p.Load().(*Person)
	fmt.Printf("p.name=%s\np.age=%v\n", _p.name, _p.age)
}
