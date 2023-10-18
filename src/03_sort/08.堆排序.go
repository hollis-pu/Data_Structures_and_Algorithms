package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 20:57
 */

func heapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐步交换根节点和末尾节点，然后重新堆化
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换根节点和最后一个节点
		heapify(arr, i, 0)              // 重新堆化，缩小堆的范围
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 寻找左右子节点中的最大值
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，交换它们并递归地堆化下去
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func main() {
	arr := []int{2, -1, 12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", arr)
	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}
