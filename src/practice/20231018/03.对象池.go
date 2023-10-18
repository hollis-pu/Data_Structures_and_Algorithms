package main

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

/**
* Description:
	这个连接池的代码写得还不是很熟练，应该多加练习！
* @Author Hollis
* @Create 2023-10-18 10:44
*/

var connChan chan *Connection

type Connection struct {
	id int
}

func InitPool(num int) {
	connChan = make(chan *Connection, num)
	for i := 1; i <= num; i++ {
		connChan <- &Connection{i}
	}
}

func GetObj(timeout time.Duration) (*Connection, error) {
	select {
	case obj := <-connChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}

}

func ReleaseObj(conn *Connection) error {
	select {
	case connChan <- conn:
		fmt.Println("release conn:", unsafe.Pointer(conn))
		return nil
	default:
		return errors.New("overflow")
	}
}

func main() {
	InitPool(10)

	for i := 0; i < 10; i++ { // 在不放回的情况下，如果取的数量超过了连接池中连接的数量，会报超时的错误。
		if v, err := GetObj(time.Second); err != nil {
			panic(err)
		} else {
			fmt.Println("get conn:", unsafe.Pointer(v))

			// 对连接进行的操作....

			if err := ReleaseObj(v); err != nil { // 操作完成后，将连接放回连接池
				panic(err)
			}
		}
	}
	fmt.Println("Done")
}
