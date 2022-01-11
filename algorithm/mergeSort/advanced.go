package main

import (
	"fmt"
	"sync"
	"time"
)

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	var s1, s2 []int
	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrency established
	go func(s *[]int) {
		defer func() { wg.Done() }()
		*s = Sort(arr[:mid])
	}(&s1)
	go func(s *[]int) {
		defer func() { wg.Done() }()
		*s = Sort(arr[mid:])
	}(&s2)
	wg.Wait()
	return merge.Merge(s1, s2)
}

func main() {
	startTime := time.Now()
	arr := []int{3, 2, 1, 6, 9, 2, 10, 54, 32, 90, 110, 1, 20, 30, 50, 17, 71, 13, 41}
	fmt.Println(Sort(arr))
	expired := time.Since(startTime)
	fmt.Println(expired)
}
