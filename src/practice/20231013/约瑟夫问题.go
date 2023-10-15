package main

import "fmt"

/**
* Description:
	成功！
* @Author Hollis
* @Create 2023-10-13 15:33
*/

type JosephusRing struct {
	id    int
	next  *JosephusRing
	first *JosephusRing
	last  *JosephusRing
}

func InitRing(num int) *JosephusRing {
	var first, last *JosephusRing
	for i := 1; i <= num; i++ {
		r := &JosephusRing{
			id: i,
		}
		if i == 1 {
			first = r
			last = r
			last.next = first
		} else {
			last.next = r
			last = r
			r.next = first
		}
	}
	return first
}
func GetRingElems(first *JosephusRing) {
	var last = first
	for {
		if last.next == first {
			fmt.Printf("id=%d--->", last.id)
			break
		}
		fmt.Printf("id=%d--->", last.id)
		last = last.next
	}
}
func PlayGame(first *JosephusRing, startNo int, countNo int) {
	var last = first
	for {
		if last.next == first {
			break
		}
		last = last.next
	}

	for i := 1; i <= startNo-1; i++ { // 移动到起始位置
		first = first.next
		last = last.next
	}

	for { // 开始游戏
		for i := 1; i <= countNo-1; i++ {
			first = first.next
			last = last.next
		}
		fmt.Printf("\n%d out", first.id)
		last.next = first.next
		first = first.next
		if first == last {
			fmt.Printf("\n%d out", first.id)
			break
		}
	}
}

func main() {
	first := InitRing(5)
	GetRingElems(first)
	PlayGame(first, 2, 3)
}
