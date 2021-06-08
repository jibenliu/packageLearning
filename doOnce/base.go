package main

import (
	"fmt"
	"sync"
)

func main()  {
	once := &sync.Once{}
	defer func() {
		if err := recover();err != nil{
			once.Do(func() {
				fmt.Println("run in recover")
			})
		}
	}()
	once.Do(func() {
		panic("panic i=0")
	})
}


// 无任何输出