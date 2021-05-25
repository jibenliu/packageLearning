package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 5)

	nums := []int{1, 2, 3, 4}

	sum(nums...)
}
