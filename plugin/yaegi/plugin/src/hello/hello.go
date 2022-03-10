package hello

import (
	"fmt"
)

func init() {
	fmt.Println("hello plugin init")
}

func CallMe(msg string) string {
	fmt.Println(msg)
	return "I am plugin"
}
