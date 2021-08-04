package main

import "fmt"

type MyInt1 int //基于类型 int 创建了新类型 MyInt1
type MyInt2 = int //创建了 int 的类型别名 MyInt2 注意类型别名的定义时 =

/**
  类型别名与类型定义的区别

 */
func main() {
	var i = 0
	var i1 MyInt1 = i //将 int 类型的变量赋值给 MyInt1 类型的变量
	//var i1 MyInt1 = MyInt1(i) //强制类型转化
	var i2 MyInt2 = i //MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	fmt.Println(i1, i2)
}
