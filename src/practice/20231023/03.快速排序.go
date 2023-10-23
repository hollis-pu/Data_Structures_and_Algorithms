package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-23 15:05
 */
func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition1(arr, left, right)
		QuickSort(arr, left, pivot)
		QuickSort(arr, pivot+1, right)
	}
}

func partition1(arr []int, left, right int) int {
	middle := left + (right-left)>>1
	for {
		for arr[left] < arr[middle] {
			left++
		}
		for arr[right] > arr[middle] {
			right--
		}
		if left >= right {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
