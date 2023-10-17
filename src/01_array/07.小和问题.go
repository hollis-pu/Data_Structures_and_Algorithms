package main

import "fmt"

/**
* Description:
	在一个数组中，一个数左边比它小的数的总和，叫数的小和所有数的小和累加起来，叫数组小和，返回数组小和。
	例子:[1, 3, 4, 2, 5]
	1左边比1小的数:没有
	3左边比3小的数:1
	4左边比4小的数:1、3
	2左边比2小的数:1
	5左边比5小的数:1、3、4、2
	所以数组小和 =1+1+3+1+1+3+4+2 =16

	这个问题可以转化为：统计i位置的右侧有多少个数字比arr[i]大。
	如：i=0是，arr[0]=1，右侧有4个数字比1大，也就是最后在计算小和时，1会被计算4次，然后i依次往后推。
	可以使用归并排序的思想来解决这个问题。
	在归并排序的过程中，当合并两个已排序的子数组时，
	可以统计右侧子数组中的每个元素比左侧子数组中的所有元素大的情况，从而计算小和。
* @Author Hollis
* @Create 2023-10-17 22:23
*/

func main() {
	arr := []int{2, 3, 5, 1, 4} // 13
	help := make([]int, len(arr))
	smallSum := calcSmallSum(arr, help, 0, len(arr)-1)
	fmt.Println(smallSum)
}

func calcSmallSum(arr, help []int, left, right int) int {
	if left >= right {
		return 0
	}
	middle := left + (right-left)>>1
	leftSmallSum := calcSmallSum(arr, help, left, middle)
	rightSmallSum := calcSmallSum(arr, help, middle+1, right)
	mergeSmallSum := merge(arr, help, left, middle, right)
	return leftSmallSum + rightSmallSum + mergeSmallSum
}

func merge(arr, help []int, left, middle, right int) int { // 左：left...middle，右：middle+1...right。help为辅助辅助，存放合并后的值。
	ans := 0
	p1 := left
	p2 := middle + 1
	i := left
	for ; p1 <= middle && p2 <= right; i++ {
		if arr[p1] < arr[p2] { // 只有左组严格小于右组时，才拷贝左右值到help数组中，并统计小和
			ans += (right - p2 + 1) * arr[p1] // 左组值乘以右组中元素的个数，添加到小和中
			help[i] = arr[p1]
			p1++
		} else { // 当左组值等于右组时，拷贝右组，并移动p2指针
			help[i] = arr[p2]
			p2++
		}
	}
	for ; p1 <= middle; i++ {
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
	return ans
}
