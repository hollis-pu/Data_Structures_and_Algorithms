package main

import (
	"fmt"
)

/**
* Description:
	稀疏数组的压缩与解压缩
* @Author Hollis
* @Create 2023-10-09 19:50
*/

var (
	originalArray [10][10]int // 定义一个二维数组
	sparseArr     []ValueNode // 定义一个切片来存放所有的值
	newArray      [10][10]int // 解压缩之后的数组
)

// ValueNode 定义一个结构体，表示存放值的结构
type ValueNode struct {
	row int
	col int
	val int
}

func main() {
	// 先给二维数组赋值
	originalArray[3][3] = 3
	originalArray[4][4] = 4
	originalArray[5][5] = 5

	// 输出查看原始数组的值
	for _, v1 := range originalArray {
		for _, v2 := range v1 {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

	compressArray(originalArray) // 转为稀疏数组
	// 遍历压缩后的值
	for _, vNode := range sparseArr {
		fmt.Println(vNode.row, vNode.col, vNode.val)
	}

	decompressArray(sparseArr) // 将压缩矩阵解压
	// 遍历解压后的数组
	for _, v1 := range newArray {
		for _, v2 := range v1 {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

}

// 将数组压缩
func compressArray(originalArray [10][10]int) {
	// 转为稀疏数组
	// 思路：
	// 遍历originalArray，如果我们发现有一个元素的值不为0，创建一个node结构体。如何设计结构体呢？
	// 将其放入到对应的切片中即可。

	valNode := ValueNode{ // 创建一个valNode 值结点，记录元素的二维数组的规模(行、列、默认值)
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode) // 将数组的结构信息也存入sparseArr中

	for i, v1 := range originalArray {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode = ValueNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
}

// 将数组还原
func decompressArray(sparseArr []ValueNode) {
	for i, v := range sparseArr {
		if i != 0 { // 第一行的值舍弃
			newArray[v.row][v.col] = v.val // 给新数组赋值
		}
	}
}
