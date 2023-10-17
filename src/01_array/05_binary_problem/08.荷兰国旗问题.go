package main

import "fmt"

/**
* Description:
	给定一个数组arr，和一个整数num，请把小于num的数放在数组的左边，等于num的数放在数组的中间，大于num的数放在数的右边。
	要求额外空间复杂度O(1)，时间复杂度O(N)。

	在这个示例中，我们使用三个指针：low、mid和high，它们分别表示数组中小于num的区域、等于num的区域和大于num的区域。
	我们从中间开始遍历数组中的每个元素，并根据其值将元素放入相应的区域。具体来说：
	如果当前元素小于num，将其与low位置的元素交换，然后将low和mid分别向前移动一位。
	如果当前元素等于num，只将mid向前移动一位。
	如果当前元素大于num，将其与high位置的元素交换，然后将high向后移动一位。
	这个过程将数组分成三个区域，分别包含小于、等于和大于num的元素，满足题目要求的条件。
	这个算法的额外空间复杂度为O(1)，时间复杂度为O(N)。
* @Author Hollis
* @Create 2023-10-17 23:12
*/

func main() {
	arr := []int{3, 1, 2, 8, 5, 2, 7, 1, 3, 6, 3}
	num := 3
	partition(arr, num)
	fmt.Println(arr)
}

func partition(arr []int, num int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] < num {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == num {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}
