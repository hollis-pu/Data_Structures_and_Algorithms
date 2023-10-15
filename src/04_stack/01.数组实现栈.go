package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-13 20:46
 */

type Stack struct {
	maxSize int    // 栈的容量
	arr     [5]int // 存放数据
	top     int    // 栈顶，初始值位-1
}

func (stack *Stack) Push(inData int) {
	if stack.top == stack.maxSize-1 {
		fmt.Println("stack is full")
		return
	}
	stack.top++
	stack.arr[stack.top] = inData
}

func (stack *Stack) List() {
	if stack.top == -1 {
		fmt.Println("stack is empty")
		return
	}
	for i := stack.top; i > -1; i-- {
		fmt.Printf("arr[%d]=%d\n", i, stack.arr[i])
	}
}

func (stack *Stack) Pop() int {
	if stack.top == -1 {
		fmt.Println("stack is empty")
		return -1
	}
	outData := stack.arr[stack.top]
	stack.top--
	fmt.Printf("%d出栈\n", outData)
	return outData
}

func main() {
	stack := &Stack{
		maxSize: 5,
		top:     -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.List()

	stack.Pop()
	stack.Pop()
	stack.List()
}
