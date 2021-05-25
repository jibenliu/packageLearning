package main

import "fmt"

//闭包 匿名函数
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()//拿到了函数体
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextIntFunc := intSeq()
	fmt.Println(nextIntFunc())
}
