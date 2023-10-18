package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 18:52
 */

type MaxHeap struct {
	data []int
	size int
}

func InitMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (heap *MaxHeap) Insert(inData int) {
	heap.size++
	heap.data = append(heap.data, inData)
	heap.heapifyUp(heap.size - 1)
}

func (heap *MaxHeap) ExtractMax() int {
	outData := heap.data[0]
	if heap.size == 0 {
		fmt.Println("heap is empty")
		return 0
	}

	heap.size--
	if heap.size > 0 {
		heap.data[0] = heap.data[heap.size]
		heap.heapifyDown(0)
	}
	return outData
}

func (heap *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if heap.data[index] > heap.data[parent] {
			heap.data[index], heap.data[parent] = heap.data[parent], heap.data[index]
		}
		index = parent
	}
}

func (heap *MaxHeap) heapifyDown(index int) {
	maxTarget := index
	for {
		left := index*2 + 1
		right := index*2 + 2
		if left <= heap.size && heap.data[left] > heap.data[maxTarget] {
			maxTarget = left
		}
		if right <= heap.size && heap.data[right] > heap.data[maxTarget] {
			maxTarget = right
		}
		if maxTarget == index {
			break
		}
		heap.data[index], heap.data[maxTarget] = heap.data[maxTarget], heap.data[index]
		index = maxTarget
	}
}

func main() {
	heap := InitMaxHeap()
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	for i := 0; i < 5; i++ {
		fmt.Println(heap.ExtractMax())
	}
}
