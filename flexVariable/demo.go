package main

// 选项模式
type Foo struct {
	key    string
	option Option
}

type Option struct {
	num int
	str string
}

type ModOption func(option *Option)

func WithNum(num int) ModOption {
	return func(option *Option) {
		option.num = num
	}
}

func WithStr(str string) ModOption {
	return func(option *Option) {
		option.str = str
	}
}

func New(key string, modOptions ...ModOption) *Foo {
	option := Option{
		num: 100,
		str: "hello",
	}

	for _, fn := range modOptions {
		fn(&option)
	}

	return &Foo{
		key:    key,
		option: option,
	}
}

func main()  {
	New("iamkey", WithNum(200), WithStr("world"))
}