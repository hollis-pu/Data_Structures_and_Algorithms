package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
  - Description:
    创建一个可以循环执行的定时任务，要求可以指定任务的次数，当达到指定执行次数时，程序退出
  - @Author Hollis
  - @Create 2023-10-18 11:06
*/
var (
	wg      sync.WaitGroup
	taskNum = 100
)

func main() {
	var timeTicker = time.NewTicker(time.Second)
	count := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer timeTicker.Stop()
		for {
			count++
			t := <-timeTicker.C
			fmt.Println(t)
			if count >= taskNum {
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("程序结束！")
}
