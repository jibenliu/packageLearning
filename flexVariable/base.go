package main

type Foo struct {
	num int
	str string
}

func New(num int, str string) *Foo {
	return &Foo{
		num: num,
		str: str,
	}
}


