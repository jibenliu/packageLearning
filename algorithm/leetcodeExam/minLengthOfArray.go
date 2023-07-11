package main

import "fmt"

func minSubArrayLen(target int, arr []int) int {
	n := len(arr)
	minLen := n + 1
	left := 0
	sum := 0

	for right := 0; right < n; right++ {
		sum += arr[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= arr[left]
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}

	return minLen
}

func main() {
	//nums := []int{2, 3, 1, 2, 4, 3}
	//target := 7
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	fmt.Println(minSubArrayLen(target, nums))
}
