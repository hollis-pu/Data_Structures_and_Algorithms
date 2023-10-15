package main

import "fmt"

/**
* Description:
	成功！
* @Author Hollis
* @Create 2023-10-13 14:59
*/

func main() {
	arr := []int{3, 6, 4, 8, 1, 2, 5, 7}
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
