package main

import "fmt"

/**
* Description:
	这段代码创建了一个 MaxHeap 结构体，其中包含最大堆的堆切片和堆大小。
	它包括了插入元素、删除最大元素、堆化（上浮和下沉）操作。
	在 main 函数中，你可以看到如何使用最大堆来插入元素并删除最大元素。最终，它会按照最大堆的性质输出元素。
* @Author Hollis
* @Create 2023-10-18 18:35
*/

type MaxHeap struct {
	data []int // 堆的元素存储在切片中
	size int   // 堆的大小
}

// NewMaxHeap 创建一个新的最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Insert 将元素插入到最大堆
func (h *MaxHeap) Insert(key int) {
	h.data = append(h.data, key) // 将新元素添加到堆的末尾
	h.size++                     // 堆大小加一
	h.heapifyUp(h.size - 1)      // 对新元素执行上浮操作，以保持最大堆性质
}

// ExtractMax 删除并返回最大元素
func (h *MaxHeap) ExtractMax() (int, error) {
	if h.size == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	max := h.data[0]         // 最大元素是根节点
	last := h.data[h.size-1] // 获取最后一个元素
	h.size--                 // 堆大小减一
	if h.size > 0 {
		h.data[0] = last // 用最后一个元素替代根节点
		h.heapifyDown(0) // 对根节点执行下沉操作，以保持最大堆性质
	}
	return max, nil
}

// heapifyUp 堆化操作，自下而上，用于插入操作
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2 // 计算父节点索引
		if h.data[index] <= h.data[parent] {
			break // 如果当前节点不大于父节点，堆性质已经满足，退出循环
		}
		// 交换当前节点和父节点的值
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent // 更新索引
	}
}

// heapifyDown 堆化操作，自上而下，用于删除操作
func (h *MaxHeap) heapifyDown(index int) {
	for {
		leftChild := 2*index + 1  // 左子节点索引
		rightChild := 2*index + 2 // 右子节点索引
		largest := index          // 假设当前节点是最大的

		// 寻找左右子节点中的最大者
		if leftChild < h.size && h.data[leftChild] > h.data[largest] {
			largest = leftChild
		}
		if rightChild < h.size && h.data[rightChild] > h.data[largest] {
			largest = rightChild
		}

		if largest == index {
			break // 如果当前节点已经是最大的，堆性质已经满足，退出循环
		}

		// 交换当前节点和最大子节点的值
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest // 更新索引（largest的位置指向的是和原来index互换值的位置，现在我们需要将index设置为该值，方便进行向下面的子树遍历）
	}
}

func main() {
	maxHeap := NewMaxHeap()

	maxHeap.Insert(10)
	maxHeap.Insert(5)
	maxHeap.Insert(20)
	maxHeap.Insert(3)
	maxHeap.Insert(30)

	fmt.Println("Max Heap:")
	for maxHeap.size > 0 {
		max, _ := maxHeap.ExtractMax()
		fmt.Printf("%d ", max)
	}
}
