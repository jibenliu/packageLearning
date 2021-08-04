package main

import "fmt"

func demo1() {
	a := "abc"
	bs := []byte(a)
	fmt.Println(bs, len(bs), cap(bs))
	// 输出： [97 98 99] 3 8
}

func demo2() {
	a := "abc"
	bs := []byte(a)
	fmt.Println(len(bs), cap(bs))
	// 输出: 3 32
}

func demo3() {
	bs := []byte("abc")
	fmt.Println(len(bs), cap(bs))
	// 输出: 3 3
}

func demo4() {
	a := ""
	bs := []byte(a)
	fmt.Println(bs, len(bs), cap(bs))
	// 输出: [] 0 0
}

func demo5() {
	a := ""
	bs := []byte(a)
	fmt.Println(len(bs), cap(bs))
	// 输出: 0 32
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}

// refer:https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651441430&idx=2&sn=9e80419fb1faa12ef7d77d72d48a2984&chksm=80bb17e4b7cc9ef2b129115938536642e780785af76b920514a4aecd29caedde630858553627&scene=21#wechat_redirect


// 字符串转字节切片步骤如下
//判断是否是常量， 如果是常量则转换为等容量等长的字节切片
//如果是变量， 先判断生成的切片是否发生变量逃逸
//如果逃逸或者字符串长度>32， 则根据字符串长度可以计算出不同的容量
//如果未逃逸且字符串长度<=32, 则字符切片容量为32


//常见逃逸情况
//函数返回局部指针
//栈空间不足逃逸
//动态类型逃逸, 很多函数参数为interface类型，比如fmt.Println(a ...interface{})，编译期间很难确定其参数的具体类型, 也会发生逃逸
//闭包引用对象逃逸