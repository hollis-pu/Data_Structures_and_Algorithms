package main

import "fmt"

/**
* Description:
	一个数组中有一个数出现了奇数次，其他数都出现了偶数次，怎么找到并打印这个数。
	我们只需要将数组中所有数进行异或运算即可。
* @Author Hollis
* @Create 2023-10-16 12:19
*/

func main() {
	arr := []int{2, 6, 4, 7, 6, 2, 7, 4, 7, 9, 9, 9} // 7 9
	result := findNums(arr)
	fmt.Println(result)
}

func findNums(arr []int) []int {
	var eor, eor1, onlyOne int
	var result = make([]int, 2)
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i] // eor=a^b
	}
	onlyOne = eor & (-eor)
	for i := 0; i < len(arr); i++ {
		if arr[i]&onlyOne != 0 {
			eor1 = eor1 ^ arr[i] // 得到a或b
		}
	}
	result[0] = eor1
	result[1] = eor ^ eor1
	return result
}
