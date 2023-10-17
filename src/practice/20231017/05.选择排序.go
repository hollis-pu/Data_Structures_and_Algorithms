package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 14:40
 */

func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	SelectionSort(arr)
	log.Println(arr)
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
