package main

/**
* Description:
* @Author Hollis
* @Create 2023-10-23 22:53
 */

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next // 先记录下翻转之前，head的下一个元素
		head.Next = pre  // 将链表翻转，head的Next指向pre
		pre = head       // 将pre和head指针后移
		head = next
	}
	return pre
}
