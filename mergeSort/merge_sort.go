package main

import (
	"fmt"
	"newTest/mergeSort/merge"
	"time"
)

//归并排序 refer: https://mp.weixin.qq.com/s/dHp-OJ0v9yNTazdYSsEKGw

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	s1 := sort(arr[:mid])
	s2 := sort(arr[mid:])
	return merge.Merge(s1, s2)
}


func main() {
	startTime := time.Now()
	arr := []int{3, 2, 1, 6, 9, 2, 10, 54, 32, 90, 110, 1, 20, 30, 50, 17, 71, 13, 41}
	fmt.Println(sort(arr))
	expired := time.Since(startTime)
	fmt.Println(expired)
}
