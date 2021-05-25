package main

import "fmt"

func main() {
	var x [3]int = [3]int{1, 2, 3}
	var y [3]int = x
	fmt.Println(x, y)
	y[0] = 999
	fmt.Println(x, y)

	var a []int = []int{1, 2, 3}
	var b []int = a

	fmt.Println(a, b)
	b[0] = 999
	fmt.Println(a, b)
}
