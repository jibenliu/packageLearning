package main

import "fmt"

func main()  {
	deferCall()
}

func deferCall()  {
	defer fmt.Println("打印前")
	defer fmt.Println("打印中")
	defer fmt.Println("打印后")
	panic("触发异常")
}
