package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	as := [...]Person{
		{Name: "小红", Age: 22},
		{Name: "小刚", Age: 18},
	}
	as[0].Age = 90
	p := Person{
		Name: "小米",
		Age:  16,
	}
	ms := map[string]Person{
		"小米": p,
	}
	//ms["小米"].Age = 111
	/**
		x = y 这种赋值的方式，你必须知道 x的地址，然后才能把值 y 赋给 x。
		当通过key获取到value时，这个value是不可寻址的
		但 go 中的 map 的 value 本身是不可寻址的，因为 map 的扩容的时候，可能要做 key/val pair迁移
		value 本身地址是会改变的
		不支持寻址的话又怎么能赋值呢
	 */

	/**
		而slice s[i] 其实是指向元素的指针，所以slice实质是通过元素的指针，修改了元素的内容，当然最终元素被修改了

		map的扩容与slice不同，那么map自身是援用类型，作为形参或返回参数的时候，传递的是值的拷贝，而值是地址，扩容时也不会扭转这个地址。而slice的扩容，会导致地址的变动。
	 */
	//fmt.Printf("%#v", ms["小米"].Name)
	s := ms["小米"]
	s.Age = 34

	/**
	字符串都是不可变的 是不可寻址的

	指针可以寻址：&Profile{}

	变量可以寻址：name := Profile{}

	字面量通通不能寻址：Profile{}
	 */
	ps := &Person{Name: "iswbm"}
	fmt.Printf("%v",*ps) // ok
}
