package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-19 10:57
 */

func main() {
	arr := []int{5, 1, 8, 3, 2, 4, -1}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	heapSize := 0
	for i := 0; i < len(arr); i++ { // 先遍历数组
		Insert(arr, i)
		heapSize++
	}
	fmt.Println(arr)
	for heapSize > 0 {
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0] // 将堆顶元素和最后一个元素做交换
		heapSize--
		heapify(arr, 0, heapSize)
	}
}

func heapify(arr []int, index, heapSize int) {
	for {
		leftChild := 2*index + 1  // 左子节点索引
		rightChild := 2*index + 2 // 右子节点索引
		largest := index          // 假设当前节点是最大的

		// 寻找左右子节点中的最大者
		if leftChild < heapSize && arr[leftChild] > arr[largest] {
			largest = leftChild
		}
		if rightChild < heapSize && arr[rightChild] > arr[largest] {
			largest = rightChild
		}

		if largest == index {
			break // 如果当前节点已经是最大的，堆性质已经满足，退出循环
		}

		// 交换当前节点和最大子节点的值
		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
	}
}

func Insert(arr []int, index int) { // 构建大根堆
	for index > 0 {
		parent := (index - 1) >> 1 // 父节点索引
		if arr[parent] < arr[index] {
			arr[parent], arr[index] = arr[index], arr[parent]
		} else {
			break
		}
		index = parent
	}
}
