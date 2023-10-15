package main

/**
* Description:
* @Author Hollis
* @Create 2023-10-12 18:20
 */

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	result := InsertSort(arr)
//	fmt.Println(result)
//}

func InsertSort(arr []int) []int {
	newArr := arr
	for i := 1; i <= len(newArr)-1; i++ {
		// 先确定第二个元素应该插入的位置
		insertVal := newArr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && newArr[insertIndex] > insertVal {
			newArr[insertIndex+1] = newArr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			newArr[insertIndex+1] = insertVal
		}
	}
	return newArr
}
