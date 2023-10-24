package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-24 9:14
 */
func main() {
	arr := []int{4, 1, 7, 2, 9, 0}
	help := make([]int, len(arr))
	MergeSort(arr, help, 0, len(arr)-1)
	fmt.Println(arr)
}

func MergeSort(arr, help []int, left, right int) {
	if left < right {
		middle := left + (right-left)>>1
		MergeSort(arr, help, left, middle)
		MergeSort(arr, help, middle+1, right)
		merge(arr, help, left, middle, right)
	}
}

func merge(arr, help []int, left, middle, right int) {
	p1 := left
	p2 := middle + 1
	i := left

	for ; p1 <= middle && p2 <= right; i++ {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}
	for ; p1 <= middle; i++ {
		help[i] = arr[p1]
		p1++
	}
	for ; p2 <= right; i++ {
		help[i] = arr[p2]
		p2++
	}
	for i := left; i <= right; i++ {
		arr[i] = help[i]
	}
}
