package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 14:44
 */
func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	QuickSort(arr, 0, len(arr)-1)
	log.Println(arr)
}

func QuickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	i := left
	j := right
	middle := left + (right-left)>>1
	for {
		for arr[i] < arr[middle] {
			i++
		}
		for arr[j] > arr[middle] {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
