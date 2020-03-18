package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("初期化処理")
}

var once sync.Once

func main() {
	// 3回呼び出しても1回しか実行されない
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
