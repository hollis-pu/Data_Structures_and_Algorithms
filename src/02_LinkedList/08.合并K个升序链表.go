package main

import "container/heap"

/**
* Description:
	[23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)
	给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。

	可以利用堆的自动有序性来实现。先将所有链表的头节点加入到堆中，构建一个小根堆（堆顶元素最小）。
	弹出堆顶元素，然后将刚才弹出元素在其链表中的下一个元素加入到堆中。然后将前一个弹出元素的Next指针指向当前弹出的元素，构成一个新的链表。
	重复上一步的过程，直到所有的链表都遍历完成，弹出所有堆中的元素，就可以得到有序的链表了。

* @Author Hollis
* @Create 2023-10-23 22:51
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode { // lists是所有链表头节点组成的切片
	if len(lists) == 0 {
		return nil
	}
	n := len(lists)
	h := &NodeHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i]) // 将每一个头节点都加入到堆中
		}
	}
	if h.Len() == 0 {
		return nil
	}
	head := heap.Pop(h).(*ListNode) // 堆中弹出一个节点作为新链表的头部（不再更改），最后返回
	pre := head                     // 往后连接链表的指针
	if pre.Next != nil {
		heap.Push(h, pre.Next) // 将弹出元素的下一个元素加入到堆中
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).(*ListNode)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return head
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
