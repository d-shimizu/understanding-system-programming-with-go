package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Ignore Ctrl + C for 10 second")

	// 可変長引数で任意の数のシグナルを設定可能
	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	time.Sleep(time.Second * 10)
}
