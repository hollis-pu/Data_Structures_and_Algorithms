package main

import "fmt"

/**
* Description:
	给定一个有序数组arr（升序）和数字num，找到小于等于num的最大index值。
* @Author Hollis
* @Create 2023-10-15 19:26
*/

func main() {
	arr := []int{1, 2, 3, 3, 3, 4, 6}
	num := 3
	result := findMaxIndex(arr, num)
	fmt.Println(result)
}

func findMaxIndex(arr []int, num int) int {
	index := -1
	left := 0
	right := len(arr) - 1
	middle := 0
	for left <= right {
		middle = left + (right-left)>>1
		if arr[middle] <= num {
			index = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return index
}
