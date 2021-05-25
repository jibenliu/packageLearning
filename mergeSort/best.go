package main

import (
	"fmt"
	"time"
)

func sort(arr []int, ch chan int) {
	defer close(ch)
	if len(arr) <= 1 {
		if len(arr) == 1 {
			ch <- arr[0]
		}
	}
	mid := len(arr) / 2
	s1 := make(chan int, mid)
	s2 := make(chan int, len(arr)-mid)

	go sort(arr[:mid], s1)
	go sort(arr[mid:], s2)
	merge(s1, s2, ch)
}

func update(s chan int, ch chan int, c *int, ok *bool) {
	ch <- *c
	*c, *ok = <-s
}

func merge(s1, s2, ch chan int) []int {
	size1 := len(s1)
	size2 := len(s2)
	finalArr := make([]int, size1+size2)
	v1, ok1 := <-s1
	v2, ok2 := <-s2
	for ok1 && ok2 {
		if v1 < v2 {
			update(s1, ch, &v1, &ok1)
		} else {
			update(s2, ch, &v2, &ok2)
		}
	}
	for ok1 {
		update(s1, ch, &v1, &ok1)
	}
	for ok2 {
		update(s2, ch, &v2, &ok2)
	}
	return finalArr
}

func main() {
	startTime := time.Now()
	ch := make(chan int)
	arr := []int{3, 2, 1, 6, 9, 2, 10, 54, 32, 90, 110, 1, 20, 30, 50, 17, 71, 13, 41}
	sort(arr, ch)
	fmt.Println(arr)
	expired := time.Since(startTime)
	fmt.Println(expired)
}
