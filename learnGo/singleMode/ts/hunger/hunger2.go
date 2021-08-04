package main

import "fmt"

type config2 struct {
	item string
}
//全局变量
var cfg *config2 = new(config2)

func NewConfig() *config2  {
	return cfg
}

func main()  {
	var c = NewConfig()
	fmt.Println(c)
}

