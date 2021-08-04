package main

import "fmt"

//非线程安全 懒汉式

//使用结构体代替类
type Tool struct {
	values int
}

//建立私有变量
var instance *Tool

func GetInstance() *Tool {
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

func main()  {
	instance := GetInstance()
	fmt.Println(instance)
}