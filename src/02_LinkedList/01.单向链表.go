package main

import "fmt"

/**
* Description:
	单向链表
* @Author Hollis
* @Create 2023-10-11 21:11
*/

type SingleLinkedList struct {
	id   int
	name string
	age  int
	next *SingleLinkedList // 指针，指向下一个元素
}

func (head *SingleLinkedList) InsertNode(newNode *SingleLinkedList) {
	tempNode := head
	for {
		if tempNode.next == nil { // 找到最后一个Node
			tempNode.next = newNode
			return
		}
		tempNode = tempNode.next
	}
}
func (head *SingleLinkedList) DeleteNode(id int) {
	exist := false
	tempNode := head
	if tempNode.next == nil { // 先判断链表是否为空，防止空指针异常
		fmt.Println("LinkedList is empty")
		return
	}

	for {
		if tempNode.next.id == id {
			exist = true
			break
		}
		tempNode = tempNode.next
		if tempNode.next == nil {
			break
		}
	}
	if exist {
		tempNode.next = tempNode.next.next
	} else {
		fmt.Printf("\nid=%d的元素不存在！\n", id)
	}
}

func (head *SingleLinkedList) GetAllNodes() {
	tempNode := head
	if tempNode.next == nil { // 先判断链表是否为空，防止空指针异常
		fmt.Println("\nLinkedList is empty")
		return
	}
	fmt.Println()
	for {
		tempNode = tempNode.next
		fmt.Printf("[%d %s %d] ---> ", tempNode.id, tempNode.name, tempNode.age)
		if tempNode.next == nil {
			return
		}
	}
}

func main() {

	head := &SingleLinkedList{} // 定义一个头结点

	elem01 := &SingleLinkedList{ // 创建一个结点
		id:   1,
		name: "Tom",
		age:  23,
	}

	elem02 := &SingleLinkedList{
		id:   2,
		name: "Jerry",
		age:  25,
	}
	elem03 := &SingleLinkedList{
		id:   3,
		name: "Tomas",
		age:  28,
	}

	head.InsertNode(elem01) // 插入结点
	head.InsertNode(elem02)
	head.InsertNode(elem03)

	head.GetAllNodes() // 获取所有结点信息
	head.DeleteNode(2) // 删除结点
	head.GetAllNodes()

}
