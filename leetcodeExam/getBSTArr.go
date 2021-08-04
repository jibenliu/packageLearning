package main

import "fmt"

/**
给你一个数组 nums 表示 1 到 n 的一个排列。
我们按照元素在 nums 中的顺序依次插入一个初始为空的二叉查找树（BST）。
请你统计将 nums 重新排序后，
统计满足如下条件的方案数：重排后得到的二叉查找树与 nums 原本数字顺序得到的二叉查找树相同。
*/

func numOfWays(nums []int) int {
	return (getAllCount(nums)%1000000007 - 1 + 1000000007) % 1000000007
}

func split(nums []int) ([]int, []int) {
	var l, r []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			l = append(l, nums[i])
		} else {
			r = append(r, nums[i])
		}
	}
	return l, r
}

func getAllCount(nums []int) int {
	if len(nums) < 1 {
		return 1
	}
	l, r := split(nums)
	s := way(len(l), len(r))
	return (s * getAllCount(l) % 1000000007) * getAllCount(r) % 1000000007
}

func way(l, r int) int {
	sum := 1
	for i := 1; i <= l; i++ {
		sum = (sum * (r + i) / i) % 1000000007
	}
	return sum
}

func main() {
	nums := []int{9, 4, 2, 1, 3, 6, 5, 7, 8, 14, 11, 10, 12, 13, 16, 15, 17, 18}
	fmt.Println(numOfWays(nums)) //216212978

	nums = []int{3, 1, 2, 5, 4, 6}
	fmt.Println(numOfWays(nums)) //19
}


/**
解题思路

1，这个题目是组合排列＋搜索树的组合题目
2，搜索树的性质，左节点<根<右节点
3，我们可以把树拆成左、根、右三部分
4，只要不改变左树内部元素的相对位置和右树内部元素的相对位置，搜索树不变
5，因此变成了排列组合问题
6，假设左树节点为m，右树为n
7,总个数为: C(len(m+n),len(m))*f(m)*f(n)
8,最后还需要把自己剪掉
 */