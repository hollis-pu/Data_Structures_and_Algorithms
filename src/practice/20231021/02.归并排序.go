package main

import "fmt"

/**
* Description:
	写得存在问题！
* @Author Hollis
* @Create 2023-10-21 17:23
*/

func main() {
	arr := []int{3, 1, 4, 8, 0, 2, 5}
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
	for ; p1 <= left && p2 <= right; i++ {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}
	for ; p1 <= left; i++ {
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
