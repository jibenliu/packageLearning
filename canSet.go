package main

import (
	"fmt"
	"reflect"
	"runtime"
)


type FooStruct struct {
	A string
	b int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v.CanSet()) // false

	if v.Elem().CanSet() {
		v.Elem().SetFloat(7.1)
	}
	fmt.Println(x)



	// struct
	a := FooStruct{}
	val := reflect.ValueOf(&a)
	fmt.Println(val.Elem().FieldByName("b").CanSet())  // false
	fmt.Println(val.Elem().FieldByName("b").CanAddr()) // true
}
