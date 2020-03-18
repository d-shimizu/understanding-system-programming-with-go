package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		// バッファ用のスライスを作成する
		buffer := make([]byte, 5)

		// 標準入力から読み込み
		size, err := os.Stdin.Read(buffer)

		// 改行コードが入力されるまでループさせる
		if err != io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d, input='%s'\n", size, string(buffer))
	}
}
