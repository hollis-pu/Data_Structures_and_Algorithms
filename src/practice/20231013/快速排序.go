package main

import "fmt"

/**
* Description:
	成功！
* @Author Hollis
* @Create 2023-10-13 15:13
*/

func main() {
	arr := []int{3, 6, 4, 8, 1, 2, 5, 7}
	result := QuickSort(arr)
	fmt.Println(result)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var left, right []int
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
