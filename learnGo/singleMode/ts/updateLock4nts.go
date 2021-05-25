package main

import "sync"

type Tools struct {
	values int
}

var newInstance *Tools
var locks sync.Locker

//双重检查
//第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了
func GetInstances() *Tools {
	if newInstance == nil {
		locks.Lock()
		if newInstance == nil {
			newInstance = new(Tools)
		}
		locks.Unlock()
	}
	return newInstance
}

//通过 sync.Once 来确保创建对象的方法只执行一次
var once sync.Once

func GetInstancesNew() *Tools {
	once.Do(func() {
		newInstance = new(Tools)
	})
	return newInstance
}
