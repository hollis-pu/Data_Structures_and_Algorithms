package main

import "fmt"

/**
* Description:
	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
	请你将两个数相加，并以相同形式返回一个表示和的链表。
	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
	https://leetcode.cn/problems/add-two-numbers/description/

	思路：先选出两个链表中较长的那个链表，然后复用这个链表。将两个链表从头节点开始，依次进行对位相加，如果有进位，则后一次进位是应该加上进位值。
	直到短的链表遍历结束，长的链表可以直接复用原来的元素。
	注意：应该注意999+1=1000这种特殊情况。
* @Author Hollis
* @Create 2023-10-23 21:01
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	node1 := &ListNode{Val: 9}
	node2 := &ListNode{Val: 9}
	node3 := &ListNode{Val: 9}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 0}
	node5 := &ListNode{Val: 0}
	node6 := &ListNode{Val: 1}
	node4.Next = node5
	node5.Next = node6
	result := addTwoNumbers(node1, node4)
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := listLength(l1)
	len2 := listLength(l2)
	var head, long, short *ListNode // head/long指向较长的链表的头节点，short指向较短的链表的头节点
	if len1 >= len2 {
		head = l1
		short = l2
	} else {
		head = l2
		short = l1
	}
	long = head      // long作为移动的指针，head作为不变量最后返回
	longLast := long // 指向长链表的最后一个元素，这个变量是为了方便最后新增节点，如：999+1=1000的最后就需要新增一个节点
	carry := 0       // 进位
	curSum := 0      // 当前相加的和

	// 整个计算分为3个阶段：1.long和short都有；2.long还有，但short已为空；3.一直到long结束了，还有进位信息。
	// 阶段1：long和short都有
	for short != nil {
		curSum = long.Val + short.Val + carry
		long.Val = curSum % 10
		carry = curSum / 10
		longLast = long
		long = long.Next
		short = short.Next
	}
	// 阶段2：short到空，但long可能不为空
	for long != nil {
		curSum = long.Val + carry
		long.Val = curSum % 10
		carry = curSum / 10
		longLast = long
		long = long.Next // 只移动long指针
	}
	// 阶段3：一直到long结束了，还有进位信息。新创建一个节点，将longLast的Next指针指向这个新节点。
	if carry != 0 {
		longLast.Next = &ListNode{Val: 1}
	}
	return head
}

func listLength(head *ListNode) int {
	var length = 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
