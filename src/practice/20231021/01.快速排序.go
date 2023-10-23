package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-21 17:13
 */
func main() {
	arr := []int{99, 3, 1, 4, 8, 0, 2, 5}
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

func partition(arr []int, left, right int) int {
	p1 := left
	p2 := right
	pivot := left + (right-left)>>1
	for {
		for arr[p1] < arr[pivot] {
			p1++
		}
		for arr[p2] > arr[pivot] {
			p2--
		}
		if p1 >= p2 {
			return p2
		}
		arr[p1], arr[p2] = arr[p2], arr[p1]
		p1++
		p2--
	}
}
