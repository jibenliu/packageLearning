package main

import (
	"fmt"
	"sync"
)

func main()  {
	once1 := &sync.Once{}
	once2 := &sync.Once{}
	once1.Do(func() {
		once2.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}

// 无任何影响，正常输出