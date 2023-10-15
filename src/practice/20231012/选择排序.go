package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 19:45
 */
func main() {
	arr := []int{5, 7, 4, 1, 9}
	SelectionSort(arr)
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		minIndex := j
		for i := j + 1; i < len(arr); i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
		}
		arr[j], arr[minIndex] = arr[minIndex], arr[j]
	}
}
