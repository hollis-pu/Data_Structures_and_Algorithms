package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-23 15:11
 */
func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
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
