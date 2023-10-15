package main

import "fmt"

/**
* Description:
	单向循环链表解决约瑟夫问题。这里和前面不同的是，头结点也是可以删除的。
	现在需要两个指针first和last，first指向第一个元素，last指向first的前一个元素（帮助我们删除指定位置的元素）。
* @Author Hollis
* @Create 2023-10-12 11:37
*/

type JosephusRing struct {
	id   int
	next *JosephusRing
}

func InitRing(num int) *JosephusRing { // 指定成员的个数进行初始化环
	first := &JosephusRing{}
	last := &JosephusRing{}
	if num < 1 {
		fmt.Println("成员数量至少为1！")
		return nil
	}
	for i := 1; i <= num; i++ {
		rData := &JosephusRing{
			id: i,
		}
		if i == 1 { // 第一个成员
			first = rData     // first指向链表的第一个元素，此后就不再改变
			last = rData      // 通过移动last来实现链表元素的增长，初始值为第一个元素
			last.next = first // 将新加入的元素指向first
		} else {
			last.next = rData // 将新的元素加入链表，并将last指向新的元素
			rData.next = first
			last = rData
		}
	}
	return first
}

func GetRingElems(first *JosephusRing) {
	if first.next == nil {
		fmt.Println("链表为空，没有任何成员！")
		return
	}
	last := first
	fmt.Println()
	for {
		fmt.Printf("[id=%d]--->", last.id)
		if last.next == first { // 遍历到最后一个元素了
			break
		}
		last = last.next
	}
}

func PlayGame(first *JosephusRing, startNo int, countNo int) {
	if first.next == nil {
		fmt.Println("链表为空！")
		return
	}
	last := first // 需要一个last指针指向最后(也就是first之前的那个元素)，因为要删除first结点，必须要找打它的前一个结点
	for {         // 将last指向最后的元素
		if last.next == first {
			break
		}
		last = last.next
	}
	// 让first和last都移动startNo-1次（移动到起始位置）（因为startNo表示的是开始的编号，但first已经指向第1个元素了，所以只需要移动startNo-1次）
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		last = last.next
	}
	// 游戏开始，开始数数（countNo）
	// first始终指向的是要删除的成员
	for {
		// 移动countNo-1次（因为成员自身也算一次）
		for i := 1; i <= countNo-1; i++ {
			first = first.next
			last = last.next
		}

		// 删除first指向的结点
		fmt.Printf("\nid=%d的元素被删除！", first.id)
		first = first.next
		last.next = first

		// 循环的终止条件：链表中只有一个元素（first=last）
		if last == first {
			break
		}
	}
	fmt.Printf("\nid=%d的元素被删除！", first.id)
}

func main() {
	first := InitRing(5)
	GetRingElems(first)
	PlayGame(first, 2, 3)
}
