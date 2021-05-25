package main

import (
	"fmt"
	"reflect"
)

type Person struct {
}

var b = func() {

}

func main() {
	a := Person{}
	fmt.Printf("a type:%T\n", a)
	fmt.Printf("a=%s\n", reflect.TypeOf(a))
	fmt.Printf("b=%s\n", reflect.TypeOf(b))

}
