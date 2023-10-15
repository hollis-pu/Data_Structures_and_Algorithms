package main

import (
	"errors"
	"fmt"
	"os"
)

/**
* Description:
	单向队列，当加入5个元素，然后再取出5个元素后，不能再添加数据。不能复用数组的队列空间。
* @Author Hollis
* @Create 2023-10-11 16:47
*/

type Queue struct {
	maxSize int
	array   [5]int // 数组 ==> 模拟队列
	front   int    // 队首（但不包含）
	rear    int    // 队尾（包含）
}

func (queue *Queue) AddQueue(val int) (err error) {

	if queue.rear == queue.maxSize-1 { // 先判断队列是否已满
		return errors.New("queue full")
	}
	queue.rear++ // rear后移
	queue.array[queue.rear] = val
	return nil
}

func (queue *Queue) GetQueue() (val int, err error) {

	if queue.rear == queue.front { // 先判断队列是否为空
		return -1, errors.New("queue empty")
	}
	queue.front++
	val = queue.array[queue.front]
	return val, err
}

func (queue *Queue) ShowQueue() {
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, queue.array[i])
	}
	fmt.Println()
}

func main() {

	// 先创建一个队列
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	// 添加数据
	var key string
	var val int

	fmt.Println("1. 输入add 表示添加数据到队列")
	fmt.Println("2. 输入get 表示从队列中获取数据")
	fmt.Println("3. 输入show 表示显示队列")
	fmt.Println("4. 输入exit 表示退出")

	for {
		fmt.Printf("请输入指令：")

		if _, err := fmt.Scanln(&key); err != nil {
			panic(err)
		}

		keyMap := map[string]struct{}{"add": struct{}{}, "get": struct{}{}, "show": struct{}{}, "exit": struct{}{}}
		if _, exist := keyMap[key]; !exist { // 输入的指令不正确
			fmt.Println("请输入正确的指令（add/get/show/exit）！")
			continue
		}

		switch key {
		case "add":
			fmt.Printf("输入需要添加到队列的数据：")
			_, err := fmt.Scanln(&val)
			if err != nil {
				panic(err)
			}

			if err := queue.AddQueue(val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功！")
			}
		case "get":
			data, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("取出元素：%d\n", data)
			}
		case "show":
			fmt.Printf("当前队列为：")
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
