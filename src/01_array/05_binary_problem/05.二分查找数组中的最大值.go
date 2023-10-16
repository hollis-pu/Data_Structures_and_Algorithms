package main

import (
	"fmt"
)

/*
*
  - Description:
    使用Master公式分析这个递归过程的时间复杂度。
    时间复杂度=2*O(N/2)+O(1) ===>a=2,b=2,d=0
    时间复杂度：O(N*log(2,2))=O(N)
  - @Author Hollis
  - @Create 2023-10-16 15:08
*/
func main() {
	arr := []int{24, 31, 60, 55, 72, 18}
	max := maxValue(arr, 0, len(arr)-1)
	fmt.Println(max)
}

func maxValue(arr []int, left int, right int) int {
	if left >= right {
		return arr[right]
	}
	middle := left + (right-left)>>1
	leftMax := maxValue(arr, left, middle)
	rightMax := maxValue(arr, middle+1, right)

	if leftMax >= rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
