package main

/**
* Description:
	前面的那种快速排序办法，消耗的时间过长，对前面算法进行改进。
	这是韩老师课上写的版本，但我个人觉得中间需要判断的条件太多了，于是求助力Chat，下一个文件是ChatGPT给出的答案，更加简洁。
* @Author Hollis
* @Create 2023-10-14 17:39
*/

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	QuickSort1(0, len(arr)-1, arr)
//	fmt.Println(arr)
//}

func QuickSort1(left, right int, arr []int) []int {
	l := left
	r := right
	pivot := arr[(l+r)/2]
	for l < r {
		for arr[l] < pivot { // 从pivot的左边找到大于等于pivot的值
			l++
		}
		for arr[r] > pivot { // 从pivot的右边找到小于等于pivot的值
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort1(left, r, arr)
	}
	if right > l {
		QuickSort1(l, right, arr)
	}

	return arr

}
