package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	one := func() {
		fmt.Println("just once")
	}

	for i := 0; i < 10; i++ {
		go func(a int) {
			once.Do(one)
		}(i)
	}
	time.Sleep(time.Millisecond*200)
}
