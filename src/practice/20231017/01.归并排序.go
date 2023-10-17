package main

import "log"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 11:20
 */
func main() {
	arr := []int{4, 2, 6, 1, 0, 8, 5, 7}
	help := make([]int, len(arr))
	MergeSort(arr, help, 0, len(arr)-1)
	log.Println(arr)
}

func MergeSort(arr []int, help []int, left int, right int) {
	if left >= right {
		return
	}
	middle := left + (right-left)>>1
	MergeSort(arr, help, left, middle)
	MergeSort(arr, help, middle+1, right)
	MergeArr(arr, help, left, middle, right)
}

func MergeArr(arr []int, help []int, left int, middle int, right int) {
	i := left
	p1 := left
	p2 := middle + 1
	for ; p1 <= middle && p2 <= right; i++ { // 防止p1和p2越界
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}
	// 遍历p1或p2中剩下的元素
	for ; p1 <= middle; i++ {
		help[i] = arr[p1]
		p1++
	}

	for ; p2 <= right; i++ {
		help[i] = arr[p2]
		p2++
	}
	for i := left; i <= right; i++ { // 将help赋值给arr（注意这里i的范围为left->right）
		arr[i] = help[i]
	}
}
