package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 12:11
 */
func main() {
	arr := []int{4, 1, 6, 8, 0, 3, 5, 7, 7, 2, 5}
	partition(arr, 5)
	fmt.Println(arr)
}

func partition(arr []int, data int) {
	left, middle, right := 0, 0, len(arr)-1
	for middle <= right { // 遍历整个数组（通过middle++和right--）
		if arr[middle] < data { // 如果当前middle位置的元素小于data，则将middle位置的元素和left位置的元素进行互换，保证left及其之前的元素都是小于data的。
			arr[left], arr[middle] = arr[middle], arr[left]
			left++ // 同时移动left和middle
			middle++
		} else if arr[middle] == data { // 如果middle位置的元素和data相等，则移动middle的位置
			middle++
		} else { // 如果当前middle位置的元素大于data，则将middle位置的元素和right位置的元素互换，保证right及其右侧的元素都大于data。然后再次循环比较当前middle的元素和data的大小关系。
			arr[middle], arr[right] = arr[right], arr[middle]
			right--
		}
	}

}
