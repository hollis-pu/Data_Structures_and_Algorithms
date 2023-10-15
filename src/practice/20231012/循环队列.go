package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 19:54
 */

type CircleQueue struct {
	maxSize int
	array   [5]int
	front   int // 指向下一个弹出元素的位置
	back    int // 指向添加元素的前一个位置
}

func (queue *CircleQueue) Push(data int) {
	queue.array[queue.back] = data
	queue.back++
}

func (queue *CircleQueue) Pop() {
	queue.front++
	fmt.Printf("删除元素：%d", queue.front)
}

func (queue *CircleQueue) GetQueue() {
	for i := queue.front + 1; i < queue.back; i++ {
		fmt.Printf("%d--->\t", queue.array[i])
	}
}

func main() {
	queue := CircleQueue{
		maxSize: 5,
		front:   -1,
		back:    0,
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Pop()
	queue.GetQueue()
}
