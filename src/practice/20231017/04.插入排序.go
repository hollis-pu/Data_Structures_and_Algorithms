package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 14:36
 */
func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	InsertSort(arr)
	log.Println(arr)
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
