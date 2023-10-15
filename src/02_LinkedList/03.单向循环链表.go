package main

import "fmt"

/**
* Description:
	单向循环链表
	这个代码中要求必须要有头结点，头结点不能被删除。不能直接用于解决约瑟夫问题（任何结点都可以被删除）
* @Author Hollis
* @Create 2023-10-12 10:44
*/

type SingleCircleLink struct {
	id   int
	next *SingleCircleLink
}

func (head *SingleCircleLink) InsertNode(newNode *SingleCircleLink) {
	if head.next == head { // 原来只有1个元素的情况
		head.next = newNode
		newNode.next = head
	} else {
		tail := head
		for {
			if tail.next == head {
				tail.next = newNode
				newNode.next = head
				break
			}
			tail = tail.next
		}
	}
}

func (head *SingleCircleLink) DeleteNode(id int) {
	if head.next == head {
		fmt.Println("当前链表只有头结点")
		return
	}
	exist := false
	tail := head
	for {
		if tail.next.id == id {
			exist = true
			break
		}
		if tail.next == head {
			fmt.Printf("\nid=%d的元素不存在！", id)
			break
		}
		tail = tail.next
	}
	if exist {
		fmt.Printf("\nid=%d的元素删除成功！", id)
		tail.next = tail.next.next
	}
}

func (head *SingleCircleLink) GetAllNodes() {
	if head.next == head { // 只有1个结点
		fmt.Printf("[id=%d] ---> ", head.id)
		return
	}
	tail := head
	for {
		if tail.next != head {
			fmt.Printf("[id=%d] ---> ", tail.id)
		} else {
			fmt.Printf("[id=%d] ---> ", tail.id) // 输出最后一个元素
			break
		}
		tail = tail.next
	}
}

func main() {
	head := &SingleCircleLink{
		id: 0,
	}

	// 初始化环形链表
	head.next = head

	node1 := &SingleCircleLink{
		id: 1,
	}
	node2 := &SingleCircleLink{
		id: 2,
	}
	node3 := &SingleCircleLink{
		id: 3,
	}

	head.InsertNode(node1)
	head.InsertNode(node2)
	head.InsertNode(node3)

	head.GetAllNodes()

	head.DeleteNode(1)
	head.DeleteNode(2)
	fmt.Println()
	head.GetAllNodes()
}
