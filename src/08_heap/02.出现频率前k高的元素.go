package main

import (
	"container/heap"
	"fmt"
)

/**
* Description:
	使用内置库container/heap实现堆
	先使用遍历数组，使用一个map来保存每个元素出现的频次，然后创建一个小根堆，小根堆的容量为k，如果超过了容量，则将堆顶元素弹出。
	使用小根堆的原因：相当于堆顶是一个门槛，如果某个元素的频次不超过堆顶，则该元素根本不需要入堆，否则加入堆中，将堆顶元素弹出，形成新的堆。
* @Author Hollis
* @Create 2023-10-23 20:24
*/

func topKFrequent(nums []int, k int) []int {
	// 创建一个 map 以存储每个元素的频次
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 创建一个最小堆
	h := &MinHeap{}
	heap.Init(h)

	// 遍历频次 map，并将元素按照频次添加到堆中
	for num, freq := range freqMap {
		heap.Push(h, Element{num, freq}) // 先将新的元素加入到堆中
		if h.Len() > k {                 // 如果堆的容量超过了k，则将堆顶元素弹出
			heap.Pop(h)
		}
	}

	// 从堆中提取前 k 个高频元素
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Element).Value
	}

	return result
}

// Element 结构用于存储元素和频次
type Element struct {
	Value int
	Freq  int
}

// MinHeap 是一个最小堆，根据频次排序
type MinHeap []Element

// 下面是实现Interface接口的方法，包括：Len()、Less()、Swap()
// Push()、Pop()是container/heap包中，Interface接口的方法，此包中Interface接口定义为：
// type Interface interface {
//	sort.Interface  // 内嵌sort.Interface接口
//	Push(x any)
//	Pop() any
//}
// 也就是说，在使用内置库container/heap时，必须手动实现：Len()、Less()、Swap()、Push()、Pop()这5个方法。

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 5, 5, 5, 6, 6, 6, 6, 1}
	k := 3
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
