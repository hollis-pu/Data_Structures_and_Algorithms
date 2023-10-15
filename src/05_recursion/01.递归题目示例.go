package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-14 16:03
 */

func testFunc(n int) {
	if n > 2 {
		n--
		testFunc(n)
	}
	fmt.Printf("n=%d\n", n)
}

func main() {
	testFunc(4)
}
