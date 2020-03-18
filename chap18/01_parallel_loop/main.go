package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"cmake...",
		"cmake . --build Release",
		"cpack",
	}

	// Wait Groupを宣言する
	var wg sync.WaitGroup

	// goroutineの開始前にsliceの長さをwait groupに登録する
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(task string) {
			// ジョブを実行する
			// このサンプルでは出力だけする
			fmt.Println(task)
			// wait gropを一つ処理する
			wg.Done()
		}(task)
	}
	wg.Wait()
}
