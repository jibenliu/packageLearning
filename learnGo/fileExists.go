package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println(Exists("sin.png"))
	fmt.Println(IsFile("sin.png"))
	fmt.Println(IsDir("sin.png"))
	fmt.Println(os.Stat("sin.png"))
	fmt.Println(os.Stat("sin1.png"))
	v, _ := os.Stat("sin.png")
	value := reflect.ValueOf(v)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Println(fmt.Sprintf("method[%d]%s and type is %v", i, typ.Method(i).Name, typ.Method(i).Type))
	}
	fmt.Println(reflect.ValueOf(v))
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
