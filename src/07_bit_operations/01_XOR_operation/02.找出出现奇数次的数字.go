package main

import "fmt"

/**
* Description:
	一个数组中有一个数出现了奇数次，其他数都出现了偶数次，怎么找到并打印这个数。
	我们只需要将数组中所有数进行异或运算即可。
* @Author Hollis
* @Create 2023-10-15 20:19
*/

func main() {
	arr := []int{2, 6, 4, 7, 6, 2, 7, 4, 7, 9, 9, 9}
	var eor int
	var result = make([]int, 2)
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i] // eor=a^b
	}
	onlyOne := eor & (-eor)
	eor1 := 0
	for i := 0; i < len(arr); i++ {
		if onlyOne&arr[i] != 0 { // 将所有数分为两类，筛选出指定位值为1的那一类
			eor1 ^= arr[i] // 找出a和b中的一个数
		}
	}
	result[0] = eor1
	result[1] = eor ^ eor1 // 求出另一个数
	fmt.Println("出现奇数次的两个数为：", result)
}
