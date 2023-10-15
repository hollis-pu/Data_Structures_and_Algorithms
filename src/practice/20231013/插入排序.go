package main

import "fmt"

/**
* Description:
	成功！
* @Author Hollis
* @Create 2023-10-13 15:25
*/

func main() {
	arr := []int{3, 6, 4, 8, 1, 2, 5, 7}
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
