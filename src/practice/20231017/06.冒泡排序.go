package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 14:42
 */

func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	BubbleSort(arr)
	log.Println(arr)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
