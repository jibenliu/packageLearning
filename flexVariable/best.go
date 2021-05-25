package main

import "fmt"

type Foo struct {
	key    string
	option Option
}

type Option struct {
	num int
	str string
}

type ModOption func(option *Option)

func New(key string, modOption ModOption) *Foo { //要求必须设置key，
	// 默认值
	option := Option{//可选参数
		num: 100,
		str: "hello",
	}
	modOption(&option)
	return &Foo{
		key:    key,
		option: option,
	}
}

func main() {
	foo := New("iamkey", func(option *Option) {
		// 调用方只设置 num
		option.num = 200
	})
	fmt.Printf("测试结构 %v", foo)
}
