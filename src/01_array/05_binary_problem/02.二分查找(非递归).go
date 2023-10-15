package main

import "fmt"

/**
* Description:
	非递归实现二分法（循环）
* @Author Hollis
* @Create 2023-10-15 16:55
*/

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		index := binarySearch1(arr, i)
		fmt.Println(index)
	}
}

func binarySearch1(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	middle := 0
	for left <= right {
		middle = left + (right-left)>>1 // 注意，middle要在for循环中更新其值
		if num == arr[middle] {
			return middle
		}
		if num < arr[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
