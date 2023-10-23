package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-23 14:54
 */
func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
	num := 3
	partition(arr, num)
	fmt.Println(arr)
}

func partition(arr []int, num int) {
	left, middle, right := 0, 0, len(arr)-1
	for middle <= right {
		if arr[middle] < num {
			arr[left], arr[middle] = arr[middle], arr[left]
			left++
			middle++
		} else if arr[middle] == num {
			middle++
		} else {
			arr[right], arr[middle] = arr[middle], arr[right]
			right--
		}
	}
}
