package main

import "fmt"

/**
* Description:
	力扣：[和等于 k 的最长子数组长度](https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/)
	给定一个数组nums和一个目标值k，找到累加和等于k的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。
* @Author Hollis
* @Create 2023-10-16 11:35
*/

func main() {
	nums := []int{6, 0, 1, -1, 5, -2, 3}
	k := 6
	result := maxSubArrayLen(nums, k)
	fmt.Println("最长连续子数组的长度为:", result) // 4
}

func maxSubArrayLen(nums []int, k int) int { // nums为原数组，k为要求的和
	maxLen := 0
	sum := 0
	sumMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i] // i的前缀和
		if sum == k {
			maxLen = i + 1
		} else if value, exist := sumMap[sum-k]; exist {
			maxLen = Max(maxLen, i-value)
		}
		if _, exist := sumMap[sum]; !exist { // key为sum的键不存在，则往sumArr中添加元素
			sumMap[sum] = i
		}
	}
	fmt.Println("sumArr:", sumMap)
	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
