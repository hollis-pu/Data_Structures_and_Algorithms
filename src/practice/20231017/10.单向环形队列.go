package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 15:53
 */

type CircleQueue struct {
	maxSize int
	arr     [5]int
	first   int // 指向下一个弹出元素的位置
	last    int // 指向下一个添加元素的位置
}

func (queue *CircleQueue) Push(inData int) {
	// 判断队列是否已满
	if (queue.last+1)%queue.maxSize == queue.first {
		fmt.Println("queue is full")
		return
	}
	queue.arr[queue.last] = inData
	queue.last = (queue.last + 1) % queue.maxSize
}

func (queue *CircleQueue) GetElems() {
	// 先判空
	if queue.last == queue.first {
		fmt.Println("queue is empty")
		return
	}
	// 获取当前队列的长度
	queueLen := (queue.last - queue.first + queue.maxSize) % queue.maxSize
	tempFirst := queue.first
	for i := 0; i < queueLen; i++ {
		fmt.Printf("arr[%d]=%d ", i, queue.arr[tempFirst])
		tempFirst = (tempFirst + 1) % queue.maxSize
	}
	fmt.Println()
}

func (queue *CircleQueue) Pop() int {
	// 先判空
	if queue.last == queue.first {
		fmt.Println("queue is empty")
		return -1
	}
	outData := queue.arr[queue.first]
	queue.first = (queue.first + 1) % queue.maxSize
	return outData
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		first:   0,
		last:    0,
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.GetElems()
	data := queue.Pop()
	fmt.Println(data)
	data = queue.Pop()
	fmt.Println(data)
	queue.GetElems()
	queue.Push(6)
	queue.Push(7)
	queue.Push(8)
	queue.GetElems()

}
