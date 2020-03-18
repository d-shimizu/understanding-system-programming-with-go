package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")

	// 終了を受け取るための終了関数付きコンテキスト cancel関数を呼び出せばキャンセルされる。
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知する
		cancel()
	}()
	// 終了を待つ
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
