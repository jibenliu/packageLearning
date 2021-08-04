package main

type Foo struct {
	option Option
}

type Option struct {
	num *int
	str *string
}

func New(option Option) *Foo {
	indexNum := -1
	if option.num == nil {
		//num实用默认值
		option.num = &indexNum
	} else {
		// option.num 即为调用方设置的初始值
	}
	return &Foo{option: option}
}

func main() {
	//调用方法，比较反人类
	num := 200
	str := "world"
	option := Option{
		num: &num,
		str: &str,
	}
	foo := New(option)
}
