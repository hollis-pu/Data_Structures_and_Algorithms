package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-21 21:54
 */
func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if i == 3 {
				break
			}
			fmt.Println(i, j)
		}
	}
}
