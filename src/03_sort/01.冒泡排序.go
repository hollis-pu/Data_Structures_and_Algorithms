package main

/**
* Description:
	从小到大排序。
	冒泡排序：比较相邻的两个元素，每轮找到一个最大值，然后置于最后（每轮冒一个）。比较n-1轮后，数组排好序。
* @Author Hollis
* @Create 2023-10-12 16:00
*/

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	result := BubbleSort(arr)
//	fmt.Println(result)
//}

func BubbleSort(arr []int) []int {
	newArr := arr
	for i := 0; i < len(arr)-1; i++ { // 外层循环，控制排序的轮数，每轮冒出一个无序数中的最大值，最后一个数无需再冒，所以i < len(arr)-1
		for j := 0; j < len(arr)-i-1; j++ { // 内层循环，控制实际排序的交换，已经有序的数无需再进行判断，所以j < len(arr)-i-1
			if newArr[j+1] < newArr[j] {
				newArr[j+1], newArr[j] = newArr[j], newArr[j+1] // Go语言中特有的交换方式
			}
		}
	}
	return newArr
}
