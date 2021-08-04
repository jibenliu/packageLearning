package main

import "fmt"

//饿汉式 直接创建好对象，不需要判断为空，同时也是线程安全，唯一的缺点是在导入包的同时会创建该对象，并持续占有在内存中

type config struct {

}

var cfg1 *config

func init()  {
	cfg1 = new(config)
}

func main()  {
	a:= cfg1
	fmt.Println(a)
}


