package main

/**
* Description:
	160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
* @Author Hollis
* @Create 2023-10-24 13:59
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 方式1：使用容器：HashSet  遍历其中一个链表，依次将每个节点的内存地址存到HashSet中，
//
//	然后将遍历另外一个链表，当第一次查询到节点的内存地址在HashSet中存在时，说明该节点就是第一次相交的节点。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{})
	tempHeadA, tempHeadB := headA, headB
	for tempHeadA != nil {
		nodeMap[tempHeadA] = struct{}{}
		tempHeadA = tempHeadA.Next
	}
	for tempHeadB != nil {
		if _, exist := nodeMap[tempHeadB]; exist {
			return tempHeadB
		}
		tempHeadB = tempHeadB.Next
	}
	return nil
}

// 方式2：不使用容器，利用偏移量。
// 由于链表A和链表B共用了一部分的链表，则可以通过遍历两个链表，得到LinkA和LinkB的长度，求出他们之间的差值Diff。
// 然后让长的链表先走Diff步，然后后面两个链表一起移动，直到两个结点相等为止。

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	tempHeadA, tempHeadB := headA, headB
	lenA, lenB := 0, 0
	for tempHeadA != nil {
		lenA++
		tempHeadA = tempHeadA.Next
	}

	for tempHeadB != nil {
		lenB++
		tempHeadB = tempHeadB.Next
	}
	tempHeadA, tempHeadB = headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			tempHeadA = tempHeadA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			tempHeadB = tempHeadB.Next
		}
	}
	for tempHeadA != nil {
		if tempHeadA == tempHeadB {
			return tempHeadA
		}
		tempHeadA = tempHeadA.Next
		tempHeadB = tempHeadB.Next
	}
	return nil
}
