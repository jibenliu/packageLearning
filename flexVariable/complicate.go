package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Key    string `json:"key"`
	Option Option
}

type Option struct {
	Num int    `json:"num"`
	Str string `json:"str"`
}

type ModOption func(option *Option)

func WithNum(num int) ModOption {
	return func(option *Option) {
		option.Num = num
	}
}

func WithStr(str string) ModOption {
	return func(option *Option) {
		if str != "" {
			option.Str = str
		}
	}
}

func NewFoo(key string, modOptions ...ModOption) *Foo {
	option := Option{
		Num: 100,
		Str: "hello",
	}

	for _, fn := range modOptions {
		fn(&option)
	}

	return &Foo{
		Key:    key,
		Option: option,
	}
}

type InputFoo struct {
	Key string `json:"key"`
	Num int    `json:"num"`
	Str string `json:"str"`
}

func main() {
	inputFoo := InputFoo{}
	input := "{\"key\":\"aaa\",\"num\":20}"

	err := json.Unmarshal([]byte(input), &inputFoo)
	if err != nil {
		panic(err)
	}
	num := inputFoo.Num
	//num, _ := strconv.Atoi(inputFoo.Num)
	foo := NewFoo(inputFoo.Key, WithNum(num), WithStr(inputFoo.Str))
	fmt.Printf(" struct is %v", foo)
}
