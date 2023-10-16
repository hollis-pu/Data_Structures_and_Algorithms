package main

import "fmt"

/**
* Description:
	归并排序，递归实现。
	归并排序-递归实现的基本思想：先选取数组中最中间的值作为middle，将数组分为左右两个数组。
	然后将左右两个数组分别进行归并排序，然后定义一个和原来数组容量相同的数组help，用两个指针p1和p2，
	分别指向左右两个数组的最开始元素（需要借助middle），取出二者中更小的元素存入数组help中，移动p1和p2，直到p1和p2越界，
	最后，使用help数组覆盖原来的arr数组，得到最后有序的数组。

	时间复杂度分析：2*O(N/2)+O(N) ===> a=2, b=2, d=1
	时间复杂度=O(N^1 * logN)=O(N*logN)，优于选择、冒泡、插入排序的O(N^2)。
* @Author Hollis
* @Create 2023-10-16 16:28
*/

func main() {
	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
	help := make([]int, len(arr)) // 新开辟一个数组空间，来存放左右合并的数组值
	MergeSort(arr, help, 0, len(arr)-1)
	fmt.Println(arr)
}

func MergeSort(arr, help []int, left, right int) {
	if left >= right {
		return
	}

	middle := left + (right-left)>>1      // 求中点
	MergeSort(arr, help, left, middle)    // 左部分有序
	MergeSort(arr, help, middle+1, right) // 右部分有序

	Merge(arr, help, left, middle, right) // 将左右数组合并到help数组中（middle参数用来分割左右数组）
}

func Merge(arr, help []int, left, middle, right int) { // 左：left...middle，右：middle+1...right。help为辅助辅助，存放合并后的值。
	p1 := left
	p2 := middle + 1
	i := left // 两个数组中比较出来的数字，应该放在help数组中的位置
	for ; p1 <= middle && p2 <= right; i++ {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}
	for ; p1 <= middle; i++ { // 将p1或p2中剩下的元素添加到help数组中
		help[i] = arr[p1]
		p1++
	}
	for ; p2 <= right; i++ {
		help[i] = arr[p2]
		p2++
	}
	for i := left; i <= right; i++ { // 将原数组arr用help数组覆盖
		arr[i] = help[i]
	}
}
