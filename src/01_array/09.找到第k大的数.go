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
	arr := []int{5, 1, 8, 2, 4}
	result := findKthLargest(arr, 2) // 4
	fmt.Println(result)
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left] // 如果左右索引相等，直接返回当前元素
	}

	pivotIndex := partition1(nums, left, right) // 通过分区函数获取基准元素的索引

	if k == pivotIndex {
		return nums[k] // 如果 k 等于基准索引，直接返回基准元素
	} else if k < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, k) // 如果 k 小于基准索引，在左半部分继续查找
	} else {
		return quickSelect(nums, pivotIndex+1, right, k) // 如果 k 大于基准索引，在右半部分继续查找
	}
}

func partition1(nums []int, left, right int) int {
	pivot := nums[(left+right)/2] // 选择最中间的元素作为基准
	i, j := left, right

	for {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
