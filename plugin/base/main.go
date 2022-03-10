package main

import (
	"log"
	"plugin"
)

type CustomPlugin interface {
	CallMe(name string) string
}

func main() {
	// 打开插件（并发安全）
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	// 在插件中搜索可导出的变量或函数
	sayHelloPlugin, err := p.Lookup("SayHelloPlugin")
	if err != nil {
		panic(err)
	}
	// 断言插件类型
	if sayHello, ok := sayHelloPlugin.(CustomPlugin); ok {
		log.Println(sayHello.CallMe("togettoyou"))
	}
}
