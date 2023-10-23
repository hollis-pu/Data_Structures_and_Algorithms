package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-23 15:01
 */
func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
	InsertSort(arr)
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		insertIndex := i
		insertVal := arr[i+1]
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		arr[insertIndex+1] = insertVal
	}

}
