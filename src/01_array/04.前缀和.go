package main

import "fmt"

/**
* Description:
	利用前缀和数组，快速求解指定范围的元素之和。
	数组前缀和（Prefix Sum）是一种用于处理数组的技术，它通过将数组中的元素累积起来，
	生成一个新的数组，其中每个元素表示原数组中某个位置之前的所有元素的总和。
	这个新数组通常被用来在常数时间内快速计算数组中任意子数组的和。
* @Author Hollis
* @Create 2023-10-15 16:42
*/

func main() {
	arr := []int{3, -1, 5, -3, 8, 7}
	sumArr := preSumArray(arr)

	fmt.Println(getSum(sumArr, 0, 5))
	fmt.Println(getSum(sumArr, 2, 4))
	fmt.Println(getSum(sumArr, 4, 5))
}

func preSumArray(arr []int) []int { //生成 前缀和 数组
	n := len(arr)
	sum := make([]int, n)
	sum[0] = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

// getSum 返回原始数组arr[left...right]范围上的累加和（利用 前缀和 数组sumArr加工得到）
func getSum(sumArr []int, left int, right int) int {
	if left <= 0 {
		return sumArr[right]
	}
	return sumArr[right] - sumArr[left-1]
}
