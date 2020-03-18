package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// ジョブ数を予め登録する
	wg.Add(2)

	go func() {
		// 非同期で仕事をする(1)
		fmt.Println("仕事1")
		// Doneで完了を通知する
		wg.Done()
	}()

	go func() {
		// 非同期で仕事をする(2)
		fmt.Println("仕事2")
		// Doneを通知する
		wg.Done()
	}()

	// 全ての処理が終わるのを待つ
	wg.Wait()
	fmt.Println("終了")
}
