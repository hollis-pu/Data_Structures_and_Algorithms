package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-15 10:56
 */

func main() {
	arr := []int{4, 2, 6, 1, 0, 8}
	result := QuickSort1(arr)
	fmt.Println(result)
}

func QuickSort1(arr []int) []int {
	var left, right []int
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort1(left)
	right = QuickSort1(right)
	return append(append(left, pivot), right...)
}
