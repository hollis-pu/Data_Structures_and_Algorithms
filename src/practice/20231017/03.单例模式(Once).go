package main

import (
	"fmt"
	"sync"
	"unsafe"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-17 12:36
 */

var once sync.Once        // 定义互斥锁
var instance1 *Singleton1 // 定义全局变量

type Singleton1 struct {
	id int
}

func GetInstance1() *Singleton1 {
	once.Do(func() {
		instance1 = &Singleton1{1}
	})
	return instance1
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			instance := GetInstance1()
			fmt.Printf("%x\n", unsafe.Pointer(instance))
			wg.Done()
		}()
	}
	wg.Wait()
}
