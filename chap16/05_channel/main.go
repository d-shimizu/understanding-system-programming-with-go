package main

import (
	"fmt"
)

func main() {
	fmt.Println("select sub()")

	// 終了を受け取るためのチャネルを生成する
	done := make(chan bool)
	// goroutineによる非同期処理
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知する
		done <- true
	}()
	// 終了を待つ
	<-done

	fmt.Println("all tasks are finished")
}
