package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-15 20:18
 */
func main() {
	a := 10
	b := 20

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
