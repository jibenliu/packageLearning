package main

type Foo struct {
	option Option
}

type Option struct {
	num int
	str string
}

func New(option Option) *Foo {
	return &Foo{option: option}
}
