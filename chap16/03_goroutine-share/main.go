package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func main() {
	// 引数渡し
	go sub1(10)

	// クロージャーのキャプチャ渡し
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()

	time.Sleep(time.Second)
}
