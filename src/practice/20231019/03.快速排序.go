package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-19 11:32
 */
func main() {
	arr := []int{5, 1, 8, 3, 2, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
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
