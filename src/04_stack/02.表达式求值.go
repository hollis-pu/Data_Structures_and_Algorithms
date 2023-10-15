package main

import (
	"errors"
	"fmt"
	"strconv"
)

/**
* Description:
	计算表达式的值，如：3*2+4-1
	现在只支持简单的加减乘除，不能加括号，运算的数字可以是多位的，如：33*90-45
* @Author Hollis
* @Create 2023-10-14 11:22
*/

type CalcStack struct {
	maxSize int
	arr     [20]int
	top     int
}

// Push 入栈
func (stack *CalcStack) Push(inData int) {
	if stack.top+1 == stack.maxSize {
		fmt.Println("stack is full")
		return
	}
	stack.top++
	stack.arr[stack.top] = inData
}

// Pop 出栈
func (stack *CalcStack) Pop() (outData int, err error) {
	if stack.top == -1 {
		return -1, errors.New("stack is empty")
	}
	outData = stack.arr[stack.top]
	stack.top--
	return
}

// IsNum 判断传入的byte是否为数字
func IsNum(s byte) bool {
	return 0 <= (int(s)-int(byte('0'))) &&
		(int(s)-int(byte('0'))) <= 9
}

// GetPriority 获取运算符的优先级，乘除优先级值为1，加减优先级值为0
func GetPriority(s byte) int {
	if string(s) == "*" || string(s) == "/" {
		return 1
	} else {
		return 0
	}
}

// Calc 传入要计算的两个数和操作符，进行运算
func Calc(num1 int, num2 int, oper int) int {

	switch string(byte(oper)) {
	case "+":
		return num2 + num1
	case "-":
		return num2 - num1
	case "*":
		return num2 * num1
	case "/":
		return num2 / num1
	default:
		panic("操作符有误！")
	}
}
func main() {
	var tempStr string
	var num1, num2, oper, result int

	str := "30+4*7-83*5" // -357
	realResult := 30 + 4*7 - 83*5

	NumStack := &CalcStack{ // 定义两个栈，一个存放数字，一个存放运算符
		maxSize: 20,
		top:     -1,
	}

	OperStack := &CalcStack{
		maxSize: 20,
		top:     -1,
	}

	for i := 0; i < len(str); i++ {
		if IsNum(str[i]) {
			if tempStr == "" {
				tempStr += string(str[i]) // 初始化tempStr
			}
			if i == len(str)-1 { // 当找到表达式的最后一位时，不再往后探寻，直接将tempStr转为整数并加入数栈
				intData, _ := strconv.Atoi(tempStr)
				NumStack.Push(intData)
				tempStr = ""
			} else {
				if IsNum(str[i+1]) { // 探寻后面一位是否是数字
					tempStr += string(str[i+1])
				} else { // 下一位是字符，则将tempStr转换为整数，然后加入数栈
					intData, _ := strconv.Atoi(tempStr)
					NumStack.Push(intData)
					tempStr = ""
				}
			}
		} else {
			if OperStack.top == -1 { // 符号栈为空，直接入栈
				OperStack.Push(int(str[i]))
			} else { // 符号栈不为空，判断新入栈符号的优先级
				if GetPriority(str[i]) > GetPriority(byte(OperStack.arr[OperStack.top])) {
					OperStack.Push(int(str[i])) // 新运算符的优先级高，直接入栈
				} else { // 新运算符的优先级低，数栈中弹出两个元素进行运算
					num1, _ = NumStack.Pop()
					num2, _ = NumStack.Pop()
					oper, _ = OperStack.Pop()
					result = Calc(num1, num2, oper)
					NumStack.Push(result) // 将结果保存到数栈中
					OperStack.Push(int(str[i]))
				}
			}
		}
	}
	// 最后计算数栈中剩余的数
	for {
		num1, _ = NumStack.Pop()
		num2, _ = NumStack.Pop()
		oper, _ = OperStack.Pop()
		result = Calc(num1, num2, oper)
		NumStack.Push(result)
		if OperStack.top == -1 { // 运算符栈为空，结束所有运算
			break
		}
	}

	fmt.Println("result:", result)
	if result == realResult {
		fmt.Println("结果正确！")
	}
}
