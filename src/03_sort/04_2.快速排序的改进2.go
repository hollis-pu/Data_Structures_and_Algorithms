package main

/**
* Description:
	ChatGPT答案版本。
* @Author Hollis
* @Create 2023-10-14 17:39
*/

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	QuickSort2(0, len(arr)-1, arr)
//	fmt.Println(arr)
//}

func QuickSort2(low, high int, arr []int) {
	if low < high {
		pivotIndex := partition(low, high, arr) // 选择分区点并获取分区点的索引
		QuickSort2(low, pivotIndex, arr)        // 递归排序分区点左侧的子数组
		QuickSort2(pivotIndex+1, high, arr)     // 递归排序分区点右侧的子数组
	}
}

// partition 将数组分成两部分，左侧的元素都小于或等于基准元素，右侧的元素都大于基准元素。
// 这是快速排序算法中的关键步骤，它在原地（in-place）完成，不需要额外的内存空间。
func partition(low, high int, arr []int) int {
	pivot := arr[(low+high)/2] // 选择最中间的元素作为基准
	i, j := low, high

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
