package main

import "fmt"

/**
* Description:
	在非递归的归并排序中，步长通常是通过一个循环结构来变化的，每一轮迭代中，步长会翻倍增加，以确保对越来越大的子数组进行合并。
	通常，初始的步长是1，然后在每轮迭代中翻倍增加。
	下面是步长如何变化的示例：
	1. 初始步骤（step = 1）：首先，将待排序的数组划分为长度为1的小数组，然后将相邻的小数组两两合并，形成长度为2的子数组。
	2. 第一轮迭代（step = 2）：在第一轮迭代中，将长度为2的子数组合并成长度为4的子数组。
	3. 第二轮迭代（step = 4）：在第二轮迭代中，将长度为4的子数组合并成长度为8的子数组。
	4. 继续迭代：这个过程会一直继续下去，每轮迭代中步长翻倍增加，直到只剩下一个长度为n的有序数组，排序完成。
	所以，步长的变化是通过简单的迭代结构来实现的，确保每轮合并的子数组大小逐渐增加，最终形成一个完全有序的数组。
* @Author Hollis
* @Create 2023-10-17 21:51
*/

func main() {
	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}

// mergeSort 使用非递归方式实现归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	n := len(arr)
	help := make([]int, n) // 辅助数组用于合并
	mergeSize := 1         // 初始化合并步长

	for mergeSize < n { // 当合并步长小于数组长度时，继续迭代
		for left := 0; left < n; left += 2 * mergeSize {
			middle := left + mergeSize - 1
			right := min(left+2*mergeSize-1, n-1) // 控制右边界不越界

			if middle < right { // 只有左侧和右侧都有元素时才进行合并
				merge(arr, help, left, middle, right)
			}
		}

		mergeSize *= 2        // 增加合并步长
		arr, help = help, arr // 交换arr和help以备下一轮合并
	}

	return arr
}

// merge 用于合并两个已排序的子数组
func merge(arr, help []int, left, middle, right int) {
	i, j, k := left, middle+1, left

	for i <= middle && j <= right {
		if arr[i] <= arr[j] {
			help[k] = arr[i]
			i++
		} else {
			help[k] = arr[j]
			j++
		}
		k++
	}

	for i <= middle {
		help[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		help[k] = arr[j]
		j++
		k++
	}
}
