package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp error")
		return
	}

	ret := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("ret = ", ret)
}
