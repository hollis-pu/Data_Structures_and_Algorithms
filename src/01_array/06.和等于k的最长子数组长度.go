package main

import "fmt"

/**
* Description:
	力扣：[和等于 k 的最长子数组长度](https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/)
	给定一个数组nums和一个目标值k，找到累加和等于k的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。
* @Author Hollis
* @Create 2023-10-15 21:55
*/

// maxSubArrayLen 函数接受一个整数数组 nums 和一个目标值 k，并返回累加和等于 k 的最长连续子数组的长度
func maxSubArrayLen(nums []int, k int) int {
	// prefixSum 是一个映射，用于存储累加和和对应的索引
	prefixSum := make(map[int]int)

	// maxLen 用于跟踪最长子数组的长度
	maxLen := 0

	// sum 用于跟踪当前的累加和
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		// 如果当前累加和等于目标值 k，整个数组就是一个候选子数组
		if sum == k {
			maxLen = i + 1
		} else if index, ok := prefixSum[sum-k]; ok {
			// 如果存在之前的累加和等于当前累加和减去 k，说明这两个位置之间的元素和为 k
			maxLen = max(maxLen, i-index)
		}

		// 如果当前累加和不在前缀和映射中，将其添加进去
		if _, ok := prefixSum[sum]; !ok {
			prefixSum[sum] = i
		}
	}

	return maxLen
}

// max 函数用于比较两个整数并返回较大的那个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, -1, 5, -2, 3}
	k := 3
	result := maxSubArrayLen(nums, k)
	fmt.Println("最长连续子数组的长度为:", result)
}
