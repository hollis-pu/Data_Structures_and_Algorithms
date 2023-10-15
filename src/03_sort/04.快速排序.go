package main

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 22:50
 */

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	result := QuickSort(arr)
//	fmt.Println(result)
//}

func QuickSort(arr []int) []int {
	var left = make([]int, len(arr))
	var right = make([]int, len(arr))
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left) // 注意，这里一定要更新left和right的值为排序后的值
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
