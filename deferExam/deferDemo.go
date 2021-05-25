package main

import "fmt"

func main() {
	var whatever [4]struct{}

	for i, j := range whatever {
		defer func() {
			fmt.Println(i)
			fmt.Println(j)
		}()
	}
	for i, j := range whatever {
		defer func(i int, j interface{}) {
			fmt.Println(i)
			fmt.Println(j)
		}(i, j)
	}
}
