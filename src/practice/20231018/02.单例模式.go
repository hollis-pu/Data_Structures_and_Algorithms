package main

import (
	"fmt"
	"sync"
	"unsafe"
)

/**
* Description:
* @Author Hollis
* @Create 2023-10-18 10:30
 */
var instance *Singleton
var mu sync.Mutex
var once sync.Once

type Singleton struct {
	id int
}

func GetInstance() *Singleton { // 互斥双重锁
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{1}
		}
	}
	return instance
}

func GetInstance1() *Singleton { // Once.Do()
	once.Do(func() {
		instance = &Singleton{1}
	})
	return instance
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
