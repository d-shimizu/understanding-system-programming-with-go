package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	// Poolを作成する。Newで新規作成時のコードを実装する
	pool := sync.Pool {
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	// 追加した要素から受け入れる
	// プールが空だと新規作成する
	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
