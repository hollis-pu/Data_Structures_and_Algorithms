package main

import "fmt"

/**
* Description:
	https://leetcode.cn/problems/palindrome-linked-list-lcci/description/

	如果不考虑空间复杂度的情况，可以使用栈，将链表的元素依次加入到栈中，然后将元素从栈中弹出，依次和原链表的元素进行比较，如果存在不相等的情况，返回false，否则就返true。
	如果限制空间复杂度为O(1)，即原地进行比较，则可以按照下面的方法进行：
	准备快慢两个指针，从链表的头节点开始移动，慢指针n1每次移动1步，快指针n2每次移动2步，这样，当n2移动到链表尾时，n1就移动到了链表的中间位置。
	然后将n1后面的链表进行反转，left指向最左侧，right指向最右侧，只需要比较left和right是否相等即可。
	注意：在返回值之前，要将链表改回最初的链表。
* @Author Hollis
* @Create 2023-10-23 21:46
*/

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3
	fmt.Println(isPalindrome(node1))

}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	n1, n2 := head, head // n1为慢指针，n2为快指针
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	// n1来到中点
	// 反转n1后面的链表，复用变量n1和n2
	n2 = n1.Next  // n2现在代表的是n1的下一个节点
	n1.Next = nil // 将n1指向nil
	var n3 *ListNode
	for n2 != nil {
		n3 = n2.Next // 在反转前，先找到n2的下一个节点，并保存到n3中
		n2.Next = n1 // 将n2向前指向，反转链表
		n1 = n2      // n1和n2都向后移
		n2 = n3
	}
	n3 = n1   // 现在的n1指向链表最右侧的位置
	n2 = head // n2指向链表最左侧的位置
	ans := true
	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			ans = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	// 将链表反转回最开始的样子
	n1 = n3.Next
	n3.Next = nil
	for n1 != nil {
		n2 = n1.Next
		n1.Next = n3
		n3 = n1
		n1 = n2
	}
	return ans
}
