package main

import "fmt"

//go中的关联数组 其他语言中的字典或hash
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("map:", m)
}
