package main

import "fmt"

/*
*
  - Description:
    给定一个数组nums和一个目标值k，找到累加和等于k的最长连续子数组长度。
    如果不存在任意一个符合要求的子数组，则返回 0。
  - @Author Hollis
  - @Create 2023-10-17 15:21
*/
func main() {
	nums := []int{2, 3, -1, -5, 4, 0, 1}
	length := maxLen(nums, 3) // 6
	fmt.Println(length)
}

func maxLen(nums []int, k int) int {
	var sum int                                  // 保存前缀和
	var maxLen int                               // 保存最大长度
	var prefixSum = make(map[int]int, len(nums)) // 前缀和数组(key为前缀和，value为index)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			maxLen = i + 1
		} else {
			if index, exist := prefixSum[sum-k]; exist {
				maxLen = maxValue(maxLen, i-index)
			}
		}
		if _, exist := prefixSum[sum]; !exist {
			prefixSum[sum] = i
		}
	}
	return maxLen
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
