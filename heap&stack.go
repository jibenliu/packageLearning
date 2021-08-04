package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func main() {
	x := foo()
	fmt.Println(*x)
}

/**
编译器会根据变量是否被外部引用来决定是否逃逸：
如果函数外部没有引用，则优先放到栈中；
如果函数外部存在引用，则必定放到堆中；

查看变量逃逸
go build -gcflags '-m -l' heap\&stack.go
加-l是为了不让foo函数被内联。


使用反汇编命令也可以看出变量是否发生逃逸。
go tool compile -S heap\&stack.go



堆上动态分配内存比栈上静态分配内存，开销大很多。
变量分配在栈上需要能在编译期确定它的作用域，否则会分配到堆上。
*/
