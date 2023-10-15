package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 19:32
 */
func main() {
	arr := []int{5, 3, 4, 1, 9}
	InsertSort(arr)
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		arr[insertIndex+1] = insertVal
	}
}
