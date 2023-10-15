package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-15 11:03
 */

func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	QuickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort2(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort2(arr, left, pivot)
		QuickSort2(arr, pivot+1, right)
	}
}
func partition(arr []int, left int, right int) int {
	pivot := arr[(left+right)/2]
	l, r := left, right
	for {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
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
