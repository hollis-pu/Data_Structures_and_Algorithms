package main

import "fmt"

/**
* Description:
	成功！
* @Author Hollis
* @Create 2023-10-13 14:33
*/

type CircleQueue struct {
	maxSize int
	array   [5]int
	front   int // 指向下一个要弹出元素的位置
	tail    int // 指向下一个加入元素的位置
}

func (queue *CircleQueue) Push(inData int) {
	if (queue.tail+1)%queue.maxSize == queue.front {
		fmt.Println("\nqueue is full")
		return
	}
	queue.array[queue.tail] = inData
	queue.tail = (queue.tail + 1) % queue.maxSize
}
func (queue *CircleQueue) Pop() (outData int) {
	if queue.tail == queue.front {
		fmt.Println("\nqueue is empty")
		return -1
	}
	outData = queue.array[queue.front]
	fmt.Printf("%d is out!\n", outData)
	queue.front = (queue.front + 1) % queue.maxSize
	return
}

func (queue *CircleQueue) GetQueue() {
	queLen := (queue.tail - queue.front + queue.maxSize) % queue.maxSize // 队列的长度
	if queLen == 0 {
		fmt.Println("queue is empty")
		return
	}
	tempFront := queue.front // 防止queue.front被修改
	for i := 1; i <= queLen; i++ {
		fmt.Printf("arr[%d]=%d\t", i, queue.array[tempFront])
		tempFront = (tempFront + 1) % queue.maxSize
	}
	fmt.Println()
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		front:   0,
		tail:    0,
	}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.GetQueue()

	queue.Pop()
	queue.Pop()
	queue.GetQueue()
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	queue.GetQueue()

}
