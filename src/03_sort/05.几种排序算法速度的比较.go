package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-14 15:27
 */
var numCount = 80000

func main() {
	var arr []int
	for i := 0; i < numCount; i++ {
		arr = append(arr, rand.Intn(numCount)) // 随机生成numCount个数字
	}

	start := time.Now().UnixMilli()
	BubbleSort(arr)
	end := time.Now().UnixMilli()
	fmt.Printf("冒泡排序耗时(ms)：%d\n", end-start)

	start = time.Now().UnixMilli()
	SelectionSort(arr)
	end = time.Now().UnixMilli()
	fmt.Printf("选择排序耗时(ms)：%d\n", end-start)

	start = time.Now().UnixMilli()
	InsertSort(arr)
	end = time.Now().UnixMilli()
	fmt.Printf("插入排序耗时(ms)：%d\n", end-start)

	start = time.Now().UnixMilli()
	QuickSort(arr)
	end = time.Now().UnixMilli()
	fmt.Printf("快速排序耗时(ms)：%d\n", end-start)

	start = time.Now().UnixMilli()
	QuickSort1(0, len(arr)-1, arr)
	end = time.Now().UnixMilli()
	fmt.Printf("快速排序1耗时(ms)：%d\n", end-start)

	start = time.Now().UnixMilli()
	QuickSort2(0, len(arr)-1, arr)
	end = time.Now().UnixMilli()
	fmt.Printf("快速排序2耗时(ms)：%d\n", end-start)
}

/*
结果：
冒泡排序耗时(ms)：9575
选择排序耗时(ms)：3458
插入排序耗时(ms)：0
快速排序耗时(ms)：56649
快速排序1耗时(ms)：4
快速排序2耗时(ms)：4
理论上来说，快速排序应该是最快的，但现在却是最慢的，而且慢得离谱，说明现在写的代码是存在问题的。
对前面快速排序的代码进行改进之后，可以看到，快速排序的时间消耗是很少的。
*/
