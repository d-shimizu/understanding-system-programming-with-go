package main

import (
	"fmt"
	"sync"
)

// グローバル変数
var id int

// メソッドの引数に*が付与されている場合は、ポインタ型の変数を引数として受け取る
func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock() をペアで呼び出してロックする
	mutex.Lock()
	defer mutex.Unlock()
	id++
//	mutex.Unlock()
	return id
}

func main() {
	// sync.Mutex構造体の変数宣言
	var mutex sync.Mutex
	// 次の宣言をしてもポインタ型になるだけで正常に動作します
	// mutex := new(sync.Mutex)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
