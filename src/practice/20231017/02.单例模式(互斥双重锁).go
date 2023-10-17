package main

import (
	"fmt"
	"sync"
)

/**
* Description:
	使用互斥双重锁
* @Author Hollis
* @Create 2023-10-17 12:08
*/

var lock sync.Mutex     // 定义互斥锁
var instance *Singleton // 定义全局变量

type Singleton struct {
	id int
}

func GetInstance() *Singleton {
	if instance == nil { // 互斥双重锁
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{1}
		}
	}
	return instance
}

func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance1 == instance2)
}
