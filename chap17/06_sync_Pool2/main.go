package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool {
		New: func() interface{} {
			count++
			return fmt.Sprintf("create: %d\n", count)
		},
	}

	pool.Put("removed: 1")
	pool.Put("removed: 2")
	// GCを呼ぶと追加された要素が消える
	runtime.GC()
	fmt.Println(pool.Get())
}
