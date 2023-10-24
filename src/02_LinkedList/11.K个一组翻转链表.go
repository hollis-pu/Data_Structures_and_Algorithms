package main

/**
* Description:
	https://leetcode.cn/problems/reverse-nodes-in-k-group/
	给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
* @Author Hollis
* @Create 2023-10-24 15:07
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 方式1：使用切片  先将链表中所有元素值存入切片中，然后在切片中进行数据的交换。(这种方式只交换了元素的值，没有交换节点)
func reverseKGroup1(head *ListNode, k int) *ListNode {
	var numArr []int
	tempHead := head
	for tempHead != nil {
		numArr = append(numArr, tempHead.Val)
		tempHead = tempHead.Next
	}
	for i := 0; i < len(numArr); i += k {
		if i+k-1 < len(numArr) {
			reverse1(numArr, i, i+k-1)
		}

	}
	tempHead = head
	for i := 0; i < len(numArr); i++ {
		tempHead.Val = numArr[i]
		tempHead = tempHead.Next
	}
	return tempHead
}
func reverse1(arr []int, left, right int) { // 包括left和right
	for i := left; i <= (left+right)>>1; i++ {
		arr[i], arr[right] = arr[right], arr[i]
		right--
	}
}

// 方式2：不使用容器
// 定义一个函数getKGroupEnd()，获取本组中最后一个节点（即当前节点之后第K-1个位置的节点，如果不足K-1个，则返回nil）
// 将本组中的每一个元素翻转，将第一组的最后一个节点（翻转前）作为整个链表的头结点（不再改变），
// 将每组的最后一个节点（翻转后）指向下一组的第一个节点，然后循环此过程。
func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := getKGroupEnd(start, k) // 从start开始，每组k个，获取最后一个节点，如果不能满足k个长度，则返回nil
	if end == nil {               // 如果第一组都不满足k的长度，直接返回head
		return head
	}

	head = end
	reverse3(start, end)
	lastEnd := start
	for lastEnd.Next != nil {
		start = lastEnd.Next
		end = getKGroupEnd(start, k)
		if end == nil {
			return head
		}
		reverse3(start, end)
		lastEnd.Next = end
		lastEnd = start
	}
	return head
}

func getKGroupEnd(start *ListNode, k int) *ListNode {
	for i := k - 1; i != 0 && start != nil; i-- {
		start = start.Next
	}
	return start
}

func reverse3(start *ListNode, end *ListNode) {
	end = end.Next
	var pre *ListNode
	var next *ListNode
	cur := start
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start.Next = end
}
