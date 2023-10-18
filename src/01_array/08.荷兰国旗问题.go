package main

import "fmt"

/**
* Description:
	给定一个数组arr，和一个整数num，请把小于num的数放在数组的左边，等于num的数放在数组的中间，大于num的数放在数的右边。
	要求额外空间复杂度O(1)，时间复杂度O(N)。
* @Author Hollis
* @Create 2023-10-18 14:10
*/

func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
	num := 3
	partition(arr, num)
	fmt.Println(arr)
}

func partition(arr []int, num int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] < num {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == num {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}
