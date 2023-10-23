package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-19 14:56
 */

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 6
	result := findKthLargest(nums, k)
	fmt.Printf("第 %d 大的数是：%d\n", k, result)
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, index int) int {
	for {
		if left > right { // 遍历完了都没找到的情况
			return 0
		}
		pivotIndex := partition2(nums, left, right)
		if pivotIndex == index {
			return nums[pivotIndex]
		}
		if pivotIndex < index {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
}

func partition2(nums []int, left, right int) int {
	l := left
	r := right
	pivot := l + (r-l)>>1
	for {
		for nums[l] < nums[pivot] {
			l++
		}
		for nums[r] > nums[pivot] {
			r--
		}
		if l >= r {
			return r
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
