package main

import "fmt"

type Student struct {
	name string
	age  int
}

func NewStudent(n string, a int) *Student {
	return &Student{name: n, age: a}
}

type StudentOptions func(*Student)

func constructStudent(opt ...StudentOptions) *Student {
	r := new(Student)
	for _, o := range opt {
		o(r)
	}
	return r
}

func setName(s string) StudentOptions {
	return func(o *Student) {
		o.name = s
	}
}

func setAge(a int) StudentOptions {
	return func(o *Student) {
		o.age = a
	}
}

func main() {
	var stu = NewStudent("张三", 20)
	fmt.Printf("stu %v", stu)
	var st = constructStudent(setName("李四"), setAge(21))
	fmt.Printf("stu %v", st)

}
