package main

import "fmt"

/**
* Description:
	递归实现二分法。返回查找值的下标，如果不存在，则返回-1。
* @Author Hollis
* @Create 2023-10-15 16:55
*/

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		index := binarySearch(arr, i, 0, len(arr)-1)
		fmt.Println(index)
	}
}

func binarySearch(arr []int, num int, left int, right int) int {
	if left > right { // 指定元素不存在时，返回-1
		return -1
	}
	//middle := (left + right) / 2
	middle := left + (right-left)>>1 // 推荐这种写法
	if num == arr[middle] {
		return middle
	}
	if num < arr[middle] {
		right = middle - 1
		return binarySearch(arr, num, left, right)
	} else {
		left = middle + 1
		return binarySearch(arr, num, left, right)
	}

}
