package main

import "fmt"

/**
桶排序概念：
	计数排序的优化版本，用有限数量的桶存放数据（存放规则由自定义函数来确定），对每个桶进行排序。
	每个桶中的数包含在一个范围，桶排序是对数据进行了区域划分，然后对每个区域内的数据进行排序，然后依次输出。
	桶排序流程描述：
	初始状态下，整个序列R[1,2,……,n]处于无序状态，且大小在[a,a+k]范围内
	设置桶的数量为bucketNum，则数据可以划分为[0,bucketNum]、[bucketNum,2*bucketNum-1]、……、[n*(bucketNum-1)/bucketNum,n]，数组中数据分别分配到相应桶中
	再对每个非空桶中的元素进行排序
	所有的非空桶依次合并即得到排序好的数据
	算法分析度过程：
	1. 时间复杂度：遍历序列O(n)，因此桶排序耗时主要取决于每个桶排序用时O(k)，总共耗时O(n+k)
	2. 稳定性：桶排序是稳定的，相同数的顺序没有改变。
refer: https://mp.weixin.qq.com/s/tzaZSKpDVbOltl1VpSxk6g
*/

func quickSort(nums []int, start, end int) []int {
	if start < end {
		i, j := start, end
		key := nums[(start+end)/2]
		for i <= j {
			for nums[i] < key {
				i++
			}
			for nums[j] > key {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if start < j {
			nums = quickSort(nums, start, j)
		}
		if end > i {
			nums = quickSort(nums, i, end)
		}
	}
	return nums
}

func bucketSort(nums []int, bucketNum int) []int {
	bucket := [][]int{}
	for i := 0; i < bucketNum; i++ {
		tmp := make([]int, 1)
		bucket = append(bucket, tmp)
	}

	//将数据分配到桶中
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]/bucketNum] = append(bucket[nums[i]/bucketNum], nums[i])
	}


	//循环所有的桶进行排序
	index := 0
	for i := 0; i < bucketNum; i++ {
		if len(bucket[i]) > 1 {
			//对每个桶内部进行排序,使用快排
			bucket[i] = quickSort(bucket[i], 0, len(bucket[i])-1)
			for j := 1; j < len(bucket[i]); j++ { //去掉一开始的tmp
				nums[index] = bucket[i][j]
				index++
			}
		}
	}
	return nums
}

func main() {
	nums := []int{45, 63, 3, 1, 29, 77, 20, 4, 30}
	fmt.Println("排序前：")
	fmt.Println(nums)
	nums = bucketSort(nums, 20)
	fmt.Println("排序后：")
	fmt.Println(nums)
}
