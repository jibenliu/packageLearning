package main

import "sync"

type student struct {
	Name string
	Age  int
}

var studentPool = &sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

func New(name string, age int) *student {
	stu := studentPool.Get().(*student)
	stu.Name = name
	stu.Age = age
	return stu
}

func Release(stu *student) {
	stu.Name = ""
	stu.Age = 0
	studentPool.Put(stu)
}

func main(){
	//stu := student.New("tom", 30)
	//defer student.Release(stu)

	stu := New("tom", 30)
	defer Release(stu)
}