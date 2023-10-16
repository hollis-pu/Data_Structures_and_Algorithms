package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-16 11:58
 */

func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	QuickSort(arr, 0, len(arr)-1)
	log.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
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
		if r <= l {
			return r
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
