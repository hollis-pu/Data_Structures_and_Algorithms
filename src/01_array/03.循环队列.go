package main

import (
	"errors"
	"fmt"
	"os"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-11 19:43
 */

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // head指向下一个弹出元素的位置
	tail    int // tail指向下一个添加元素的位置
}

// Push 向队列中添加元素
func (queue *CircleQueue) Push(data int) (err error) {
	if (queue.tail+1)%queue.maxSize == queue.head { // 队列已满
		return errors.New("queue full")
	}
	queue.array[queue.tail] = data
	queue.tail = (queue.tail + 1) % queue.maxSize
	return
}

// Pop 从队列中弹出元素
func (queue *CircleQueue) Pop() (data int, err error) {
	if queue.head == queue.tail { // 队列为空
		return -1, errors.New("queue empty")
	}
	data = queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	return data, nil
}

func (queue *CircleQueue) ShowCircleQueue() {
	queueLen := (queue.tail - queue.head + queue.maxSize) % queue.maxSize // 队列中元素个数
	if queueLen == 0 {
		fmt.Println("queue is empty")
		return
	}
	originalHead := queue.head
	for i := 1; i <= queueLen; i++ { // 遍历队列长度个元素
		fmt.Printf("array[%d]=%d\t", i, queue.array[originalHead])
		originalHead = (originalHead + 1) % queue.maxSize
	}
	fmt.Println()
}

func main() {

	// 先创建一个队列
	queue := CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	// 添加数据
	var key string
	var val int

	fmt.Println("1. 输入push 表示添加数据到队列")
	fmt.Println("2. 输入pop 表示从队列中获取数据")
	fmt.Println("3. 输入show 表示显示队列")
	fmt.Println("4. 输入exit 表示退出")

	for {
		fmt.Printf("请输入指令：")

		if _, err := fmt.Scanln(&key); err != nil {
			panic(err)
		}

		keyMap := map[string]struct{}{"pop": struct{}{}, "push": struct{}{}, "show": struct{}{}, "exit": struct{}{}}
		if _, exist := keyMap[key]; !exist { // 输入的指令不正确
			fmt.Println("请输入正确的指令（add/get/show/exit）！")
			continue
		}

		switch key {
		case "push":
			fmt.Printf("输入需要添加到队列的数据：")
			_, err := fmt.Scanln(&val)
			if err != nil {
				panic(err)
			}

			if err := queue.Push(val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功！")
			}
		case "pop":
			data, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("取出元素：%d\n", data)
			}
		case "show":
			fmt.Printf("当前队列为：")
			queue.ShowCircleQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
