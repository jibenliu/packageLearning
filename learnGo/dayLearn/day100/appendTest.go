package main

import "fmt"

func main() {
	m := map[string]int{
		"G": 7, "A": 1,
		"C": 3, "E": 5,
		"D": 4, "B": 2,
		"F": 6, "I": 9,
		"H": 8,
	}
	var order []string
	for k, _ := range m {
		order = append(order, k)
	}
	fmt.Println(order) //随机字符串
}
