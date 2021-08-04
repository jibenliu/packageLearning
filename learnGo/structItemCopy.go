package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type A struct {
	A int
	B int
	C string
	E string
	F string
}

type B struct {
	B int
	C string
	D int
	E string
	G string
}

func main() {
	//法一
	a := &A{1, 1, "a", "b", "f"}
	aj, _ := json.Marshal(a)
	b := new(B)
	_ = json.Unmarshal(aj, b)

	fmt.Printf("%+v\n", b)

	c := new(B)
	CopyStruct(a, c)
	fmt.Printf("%+v\n", c)
}

func CopyStruct(src, dst interface{}) {
	sVal := reflect.ValueOf(src).Elem()
	dVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < sVal.NumField(); i++ {
		value := sVal.Field(i)
		name := sVal.Type().Field(i).Name

		dValue := dVal.FieldByName(name)
		if dValue.IsValid() == false {
			continue
		}
		dValue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
}
