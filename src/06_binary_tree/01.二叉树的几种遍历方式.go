package main

import "fmt"

/**
* Description:
	前序遍历：根---左---右
	中序遍历：左---根---右
	后序遍历：左---右---根
* @Author Hollis
* @Create 2023-10-15 15:26
*/

type BinaryTree struct {
	id    int
	left  *BinaryTree
	right *BinaryTree
}

func InitTree() *BinaryTree {
	node1 := &BinaryTree{
		id: 1,
	}
	node2 := &BinaryTree{
		id: 2,
	}
	node3 := &BinaryTree{
		id: 3,
	}
	node4 := &BinaryTree{
		id: 4,
	}
	node5 := &BinaryTree{
		id: 5,
	}
	node6 := &BinaryTree{
		id: 6,
	}
	node7 := &BinaryTree{
		id: 7,
	}

	node4.left = node2 // node4为根节点
	node4.right = node6
	node2.left = node1
	node2.right = node3
	node6.left = node5
	node6.right = node7

	return node4
}

func PreOrder(node *BinaryTree) { // 前序遍历
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.id)
	PreOrder(node.left)
	PreOrder(node.right)
}

func InOrder(node *BinaryTree) { // 中序遍历
	if node == nil {
		return
	}
	InOrder(node.left)
	fmt.Printf("%d ", node.id)
	InOrder(node.right)
}

func PostOrder(node *BinaryTree) { // 后序遍历
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	fmt.Printf("%d ", node.id)
}

func main() {
	root := InitTree()
	fmt.Printf("\nPreorder traversal: ")
	PreOrder(root)
	fmt.Printf("\nInorder traversal: ")
	InOrder(root)
	fmt.Printf("\nPostorder traversal: ")
	PostOrder(root)
}
