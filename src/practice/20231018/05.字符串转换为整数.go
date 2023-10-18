package main

import (
	"errors"
	"fmt"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 11:34
 */

func main() {
	str := "-12345"
	result, err := Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Atoi(s string) (int64, error) {
	newStr := s
	var isNegative int64 = 1
	if s[0] == '-' {
		isNegative = -1
		newStr = s[1:]
	}

	var sum int64
	for _, runeData := range newStr {
		num := runeData - '0'
		if num <= 0 || num >= 9 {
			return 0, errors.New("字符串有误！")
		}
		sum = sum*10 + int64(num)
	}
	return isNegative * sum, nil
}
