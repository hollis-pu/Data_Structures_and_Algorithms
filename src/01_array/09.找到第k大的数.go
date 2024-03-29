package main

import "fmt"

/*
*
  - Description:
    给你一个整数数组nums和一个正数k，请你返回第k大的数。要求时间复杂度O(n)。
    可以把这个问题转化为：确定排完序(升序)后，第 len(num)-k 位置上的数，这个数就是我们的最后结果。
    但我们不需要真的去将数组nums排好序，可以借鉴快速排序中的partition函数（左边的数比pivot小，右边的数比pivot大）的思想。
    如果pivot刚好等于了我们的目标位置 len(num)-k，则说明这个位置上的数就是我们要找的数。
    如果pivot小于目标位置 len(num)-k，说明我们要找的数在pivot的右侧，反之则在左侧。（只需要找一侧就可以了，这样就提高了查找的效率）
  - @Author Hollis
  - @Create 2023-10-18 16:38
*/
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 1
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
