package main

import "fmt"

func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func func1() (r int) {
	// 1.赋值
	r = 0
	// 2.闭包引用，返回值被修改
	defer func() {
		r++
	}()
	// 3.空的 return
	return
}

func func2() (r int) {
	t := 5
	// 1.赋值
	r = t
	// 2.闭包引用，但是没有修改返回值 r
	defer func() {
		t = t + 5
	}()
	// 3.空的 return
	return
}

func func3() (r int) {
	// 1.赋值
	r = 1
	// 2.r 作为函数参数，不会修改要返回的那个 r 值
	defer func(r int) {
		r = r + 5
	}(r)
	// 3.空的 return
	return
}

func main() {
	/**
	return xxx
	这条语句并不是一个原子指令，经过编译之后，变成了三条指令：

	1. 返回值 = xxx
	2. 调用 defer 函数
	3. 空的 return
	*/
	result := watShadowDefer(50)
	fmt.Println(result) //100
	r1 := func1()
	fmt.Println(r1) //1

	r2 := func2()
	fmt.Println(r2) //5

	r3 := func3()
	fmt.Println(r3) //1
}
