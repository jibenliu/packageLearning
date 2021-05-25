package main

import "fmt"

func main() {
	var data = [10]byte{'T', 'S', 'E'}
	var str = string(data[:])
	fmt.Println(str)
}
