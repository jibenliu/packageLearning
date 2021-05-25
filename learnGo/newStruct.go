package main

import (
	"fmt"
	"unsafe"
)

type Person struct{
	name string
	age uint8
}

func NewPerson(name string, age uint8) *Person{
	return &Person{name, age}
}

func main(){
	p := NewPerson("Jack", 20)
	fmt.Println("%+v\n", p)
	fmt.Println(unsafe.Sizeof(Person{}))
}