package main

import "fmt"

/*
*
  - Description:
    失败！
  - @Author Hollis
  - @Create 2023-10-23 15:37
*/
func main() {
	arr := []int{0, 1, 3, 4, 5, -1, 2}
	result := FindKthMax(arr, 3)
	fmt.Println(result)
	//result := partition2(arr, 0, len(arr)-1)
	//fmt.Println(arr)
	//fmt.Println(result)
}

func FindKthMax(arr []int, k int) int {
	return findByIndex(arr, 0, len(arr)-1, len(arr)-k)
}

func findByIndex(arr []int, left, right, index int) int {
	if left == right {
		return arr[left]
	}
	returnIndex := partition2(arr, left, right)
	if index == returnIndex {
		return arr[index]
	} else if index < returnIndex {
		return findByIndex(arr, left, returnIndex-1, index)
	} else {
		return findByIndex(arr, returnIndex+1, right, index)
	}
}

func partition2(arr []int, left, right int) int {
	middle := left + (right-left)>>1
	for {
		for arr[left] < arr[middle] {
			left++
		}
		for arr[right] > arr[middle] {
			right--
		}
		if left >= right {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
