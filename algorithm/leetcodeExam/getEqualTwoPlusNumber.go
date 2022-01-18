package main

import (
	"fmt"
	"time"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9 所以返回 [0, 1]

// 暴力法
func Rude(nums []int, target int) []int {
	length := len(nums)
	ans := make([]int, 0)
Loop:
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
				break Loop
			}
		}
	}
	return ans
}

type Answer struct {
	P [][2]int
}

// 获取所有的数字对
func Rude2(nums []int, target int) *Answer {
	length := len(nums)
	ans := new(Answer)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				ans.P = append(ans.P, [2]int{i, j})
			}
		}
	}
	return ans
}

// 哈希表法
// 通过以空间换取速度的方式，我们可以将查找时间从 O(n) 降低到 O(1)。哈希表正是为此目的而构建的，它支持以近似恒定的时间进行快速查找。
// 我们可以将每个元素的值和它的索引添加到哈希表中，然后在遍历时检查每个元素对应的目标元素（target - nums[i]）是否存在于表中。注意⚠️，该目标元素不能是 nums[i] 本身！

func twoSum(nums []int, target int) []int {
	length := len(nums)
	result := make([]int, 0)
	m := make(map[int]int, length)
	for i, k := range nums {
		// 判断map中是否存在key为[target - k]的值
		if value, exist := m[target-k]; exist && value != i {
			// append：尾部追加元素
			result = append(result, value)
			result = append(result, i)
			break
		}
		m[k] = i
	}
	return result
}

func main() {
	nums := []int{1, 10, 3, 5, 2, 7, 11, 15}
	target := 9
	t1 := time.Now()
	ret := twoSum(nums, target)
	fmt.Println(ret)
	fmt.Println(time.Since(t1))

	t2 := time.Now()
	res := Rude(nums, target)
	fmt.Println(res)
	fmt.Println(time.Since(t2))
}
