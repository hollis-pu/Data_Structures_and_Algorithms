package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-21 17:40
 */

type Heap struct {
	arr  []int
	size int
}

func main() {
	arr := []int{99, 3, 1, 4, 8, 0, 2, 5}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	heap := Heap{}
	for i := 0; i < len(arr); i++ {
		heap.arr[heap.size] = arr[i]
		heap.size++
		Insert(&heap, heap.size-1)
	}
}

func Insert(heap *Heap, inData int) {
	curIndex := heap.size - 1
	if curIndex == 0 {
		heap.arr[0] = inData
	} else {
		for {
			parent := (curIndex - 1) / 2
			if inData > heap.arr[parent] {
				inData, heap.arr[parent] = heap.arr[parent], inData
			}
		}

	}
}
