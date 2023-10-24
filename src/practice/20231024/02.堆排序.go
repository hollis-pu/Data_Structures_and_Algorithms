package main

import (
	"container/heap"
	"fmt"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-24 9:59
 */

func main() {
	arr := []int{5, 1, 6, 0, 3, 4}
	result := HeapSort(arr)
	fmt.Println(result)
}

func HeapSort(arr []int) []int {
	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < len(arr); i++ {
		heap.Push(h, Element{arr[i]})
	}

	for i := 0; h.Len() > 0; i++ {
		arr[i] = heap.Pop(h).(Element).value
	}
	return arr
}

type Element struct {
	value int
}

type minHeap []Element

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
