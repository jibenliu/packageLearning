package main

import (
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	// 初始化解释器
	i := interp.New(interp.Options{GoPath: "./plugin/"})

	// 插件需要使用标准库
	if err := i.Use(stdlib.Symbols); err != nil {
		panic(err)
	}

	// 导入 hello 包
	if _, err := i.Eval(`import "hello"`); err != nil {
		panic(err)
	}

	// 调用 hello.CallMe
	v, err := i.Eval("hello.CallMe")
	if err != nil {
		panic(err)
	}
	callMe := v.Interface().(func(string) string)
	fmt.Println(callMe("togettoyou"))
}
