package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 12:11
 */
func main() {
	arr := []int{4, 1, 6, 8, 0, 3, 5, 7, 7, 2, 5}
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
	l := left
	r := right
	pivot := l + (r-l)>>1
	for {
		for arr[l] < arr[pivot] {
			l++
		}
		for arr[r] > arr[pivot] {
			r--
		}
		if l >= r {
			return r
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
