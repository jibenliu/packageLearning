package main

import (
	"fmt"
	"sync"
)

func main()  {
	once := &sync.Once{}
	once.Do(func() {
		once.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}

// 死锁，第二个do等待第一个释放锁，第一个被第二个阻塞了