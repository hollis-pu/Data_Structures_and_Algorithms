package main

/**
* Description:
	[142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)
* @Author Hollis
* @Create 2023-10-24 14:36
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 方式1：使用HashSet
func detectCycle(head *ListNode) *ListNode {
	tempHead := head
	nodeMap := make(map[*ListNode]struct{})
	for tempHead != nil {
		if _, exist := nodeMap[tempHead]; exist {
			return tempHead
		} else {
			nodeMap[tempHead] = struct{}{}
		}
		tempHead = tempHead.Next
	}
	return nil
}

// 方式2：不使用容器：快慢指针相遇
// 这分为两个过程：1.设置快慢两个指针，慢指针s每次走1步，慢指针f每次走两步。如果链表有环，则s和f一定会在某个位置相遇。
// 2.将f重新设置为链表的头节点，s还是在相遇的位置。现在让s和f都每次走一步，s和f再次相遇的位置就是第一次出现环的位置。（这是一个结论，可以直接记住）
func detectCycle1(head *ListNode) *ListNode {
	s, f := head, head
	isCycle := false
	for f != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			isCycle = true
			break
		}
	}
	if isCycle {
		return nil
	}
	f = head
	for {
		if s == f {
			return s
		}
		s = s.Next
		f = f.Next
	}
}
