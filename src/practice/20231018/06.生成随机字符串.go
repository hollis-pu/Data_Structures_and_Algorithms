package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 11:51
 */
func main() {
	result := GetRandomString(4)
	fmt.Println(result)
}

func GetRandomString(n int) string {
	var result string
	characterSet := "abcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 注意，指定随机数种子应该在for循环之外，否则生成的随机数基本是同一个数。
	for i := 0; i < n; i++ {
		index := r.Intn(len(characterSet) - 1)
		result += string(characterSet[index])
	}
	return result
}
