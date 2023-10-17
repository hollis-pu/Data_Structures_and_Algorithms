package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 15:09
 */

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearch(arr, 0)
	fmt.Println(index)
}

func BinarySearch(arr []int, data int) int {
	left := 0
	right := len(arr) - 1
	for {
		middle := left + (right-left)>>1
		if data == arr[middle] {
			return middle
		}
		if data < arr[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
		if left > right {
			return -1
		}
	}
}
