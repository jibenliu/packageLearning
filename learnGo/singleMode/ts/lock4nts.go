package main

import (
	"fmt"
	"sync"
)

//非线程安全基础上加锁保证线程安全

//使用结构体代替类
type Tool struct {
	values int
}

//建立私有变量
var instance *Tool
var lock sync.Locker

func GetInstance() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

func main() {
	instance := GetInstance()
	fmt.Println(instance)
}
