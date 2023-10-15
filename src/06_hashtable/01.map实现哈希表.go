package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-15 12:08
 */

func main() {
	// 创建一个空的哈希表
	hashMap := make(map[string]int)

	// 插入键值对
	hashMap["apple"] = 1
	hashMap["banana"] = 2
	hashMap["cherry"] = 3

	// 查找键的值
	fmt.Println("banana:", hashMap["banana"]) // 输出: banana: 2

	// 删除键值对
	delete(hashMap, "banana")

	// 检查键是否存在
	_, exists := hashMap["banana"]
	fmt.Println("Is banana in the hash map?", exists) // 输出: Is banana in the hash map? false
}
