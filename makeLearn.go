package main

import (
	"fmt"
)

func main() {
	a := make([]func() int, 3)
	a[0] = func() int {
		return 1
	}
	a[1] = func() int {
		return 2
	}
	a[2] = func() int {
		return 3
	}
	for index, funcName := range a {
		fmt.Println(index)
		funcName()
	}
	b := []func() int{func() int {
		return 1
	}, func() int {
		return 2
	}, func() int {
		return 3
	}}
	for index, funcName := range b {
		fmt.Println(index)
		funcName()
	}

	//c := new([]func() int)
	//c[0] = func() int {
	//	return 1
	//}
	//c[1] = func() int {
	//	return 2
	//}
	//c[2] = func() int {
	//	return 3
	//}
	//
	//for index, funcName := range c {
	//	fmt.Print(index)
	//	funcName()
	//}

	e := Black
	test(e)
	//􏰆􏴊􏴋􏵒􏰆􏴊􏴋􏵒不支持隐式类型转换
	//x := 1
	//test(x) //Error: cannot use x (type int) as type Color in function argument
	test(1) //显式转换

	s := []int{1, 2, 3}
	println(newTest("sum: %d", s...))
}

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {}

func newTest(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
