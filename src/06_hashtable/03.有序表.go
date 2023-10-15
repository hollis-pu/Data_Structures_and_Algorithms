package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

/**
* Description:
	使用第三方库，实现有序表
* @Author Hollis
* @Create 2023-10-15 22:09
*/

func main() {
	m := treemap.NewWithIntComparator() // empty (keys are of type int)
	m.Put(1, "x")                       // 1->x
	m.Put(2, "b")                       // 1->x, 2->b (in order)
	m.Put(1, "a")                       // 1->a, 2->b (in order)
	fmt.Println(m)                      // map[1:a 2:b]
	_, _ = m.Get(2)                     // b, true
	_, _ = m.Get(3)                     // nil, false
	_ = m.Values()                      // []interface {}{"a", "b"} (in order)
	_ = m.Keys()                        // []interface {}{1, 2} (in order)
	m.Remove(1)                         // 2->b
	m.Clear()                           // empty
	m.Empty()                           // true
	m.Size()                            // 0

	// Other:
	m.Min() // Returns the minimum key and its value from map.
	m.Max() // Returns the maximum key and its value from map.
}
