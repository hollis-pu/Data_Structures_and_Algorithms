package main

/**
* Description:
	从小到大排序
	选择排序：每轮选择一个无需中的最小值，n-1轮后，数组有序。
	技巧：先写内层循环，然后再嵌套一个外层循环。
	先假设我们只需要找一次最小值，具体操作如下：
		1.假设arr[0]就是最小元素(min=arr[0])；（注意：需要使用一个变量index记录下标0，后面交换的时候要用到）
		2.遍历数组的每个元素，依次与min进行比较，如果有比min值更小的，则更新min的值和下标index；
		3.交换arr[0]和arr[index]的值。

	这样就完成了一次最小值的排序，我们只需要在外面嵌套一层for循环，将0改成循环变量j即可。这样就大大简化了我们写代码的过程。

* @Author Hollis
* @Create 2023-10-12 16:15
*/

//func main() {
//	arr := []int{3, 8, 1, 5, 4, 7, 6, 9, 2}
//	result := SelectionSort(arr)
//	fmt.Println(result)
//}

func SelectionSort(arr []int) []int {
	newArr := arr
	for j := 0; j < len(arr)-1; j++ {
		//minArr := newArr[j] // 之前用了一个变量来存最小值，后面发现可以优化一下，只需要记录下标minIndex就可以了，所以将这里注释掉
		minIndex := j
		for i := j + 1; i < len(arr); i++ { // 找出最小值的下标
			if newArr[i] < newArr[minIndex] {
				minIndex = i // 更新minIndex
			}
		}
		newArr[minIndex], newArr[j] = newArr[j], newArr[minIndex]
	}
	return newArr
}
