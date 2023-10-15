package main

import "fmt"

/*
*
  - Description:
    给定一个有序数组arr（升序）和数字num，找到大于等于num的最小index值。
  - @Author Hollis
  - @Create 2023-10-15 19:16
*/
func main() {
	arr := []int{1, 2, 3, 2, 3, 4, 2}
	num := 2
	result := findMinIndex(arr, num)
	fmt.Println(result)
}

func findMinIndex(arr []int, num int) int {
	index := -1
	left := 0
	right := len(arr) - 1
	middle := 0
	for left <= right {
		middle = left + (right-left)>>1
		if arr[middle] >= num {
			index = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return index
}
