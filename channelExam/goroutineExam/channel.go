package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan string, 512)
var arr1 = []string{"qq", "ww", "ee", "rr", "tt"}

func chanTest1() {
	for _, v := range arr1 {
		chan1 <- v
	}
	close(chan1)
}

func chanTest2() {
	for {
		getStr, ok := <-chan1 // 阻塞,直到chan1里面有数据
		if !ok {              // 判断channel是否关闭或者为空
			return
		}
		fmt.Println(getStr) // 按数组顺序内容输出
	}
}

func main() {
	go chanTest1()
	go chanTest2()
	time.Sleep(time.Microsecond * 200)
}
