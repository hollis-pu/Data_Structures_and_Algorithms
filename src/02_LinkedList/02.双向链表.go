package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-11 22:16
 */

type DoubleLinkedList struct {
	id       int
	name     string
	age      int
	previous *DoubleLinkedList // 指向前一个元素
	next     *DoubleLinkedList // 指针，指向下一个元素
}

func (head *DoubleLinkedList) InsertNode(newNode *DoubleLinkedList) {
	tempNode := head
	for {
		if tempNode.next == nil { // 找到最后一个Node
			tempNode.next = newNode
			newNode.previous = tempNode
			return
		}
		tempNode = tempNode.next
	}
}
func (head *DoubleLinkedList) DeleteNode(id int) {
	exist := false
	tempNode := head
	if tempNode.next == nil { // 先判断链表是否为空，防止空指针异常
		fmt.Println("DoubleLinkedList is empty")
		return
	}

	for {
		tempNode = tempNode.next
		if tempNode.id == id {
			exist = true
			break
		}
		if tempNode.next == nil {
			break
		}
	}
	if exist {
		tempNode.previous.next = tempNode.next
		tempNode.next.previous = tempNode.previous
		//tempNode.next = nil  // 2023.10.11 这里到底需不需要写呢？
		//tempNode.previous = nil
	} else {
		fmt.Printf("\nid=%d的元素不存在！\n", id)
	}
}

func (head *DoubleLinkedList) GetAllNodes() {
	tempNode := head
	if tempNode.next == nil { // 先判断链表是否为空，防止空指针异常
		fmt.Println("\nDoubleLinkedList is empty")
		return
	}
	fmt.Println()
	for {
		tempNode = tempNode.next
		fmt.Printf("[%d %s %d] <---> ", tempNode.id, tempNode.name, tempNode.age)
		if tempNode.next == nil {
			return
		}
	}
}

func (head *DoubleLinkedList) GetReverseNodes() {
	tempNode := head
	if tempNode.next == nil { // 先判断链表是否为空，防止空指针异常
		fmt.Println("\nDoubleLinkedList is empty")
		return
	}
	fmt.Println()
	for { // 找到最后一个Node
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}
	for { // 向前遍历
		if tempNode.previous == nil {
			break
		}
		fmt.Printf("[%d %s %d] <---> ", tempNode.id, tempNode.name, tempNode.age)
		tempNode = tempNode.previous
	}
}
func main() {

	head := &DoubleLinkedList{} // 定义一个头结点

	elem01 := &DoubleLinkedList{ // 创建一个结点
		id:   1,
		name: "Tom",
		age:  23,
	}

	elem02 := &DoubleLinkedList{
		id:   2,
		name: "Jerry",
		age:  25,
	}
	elem03 := &DoubleLinkedList{
		id:   3,
		name: "Tomas",
		age:  28,
	}

	head.InsertNode(elem01) // 插入结点
	head.InsertNode(elem02)
	head.InsertNode(elem03)

	head.GetAllNodes()     // 获取所有结点信息
	head.GetReverseNodes() // 向前遍历结点

	head.DeleteNode(2) // 删除结点
	head.GetAllNodes()
	head.GetReverseNodes() // 向前遍历结点
}
