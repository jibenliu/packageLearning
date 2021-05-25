package main

import (
	"fmt"
)

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}
func main() {
	printType(1024)
	printType("pig")
	printType(true)
}
