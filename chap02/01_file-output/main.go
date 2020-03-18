package main

import (
	"os"
)

func main() {
	// ファイル生成
	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	// ファイルへの書き込み
	file.Write([]byte("os.File example\n"))
	file.Close()
}

