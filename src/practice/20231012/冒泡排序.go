package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 19:50
 */
func main() {
	arr := []int{5, 7, 4, 1, 9}
	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
