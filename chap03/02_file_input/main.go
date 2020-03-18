package main

import (
	"io"
	"os"
)

func main() {
	// ファイルを読み込む
	file, err := os.Open("./01_stdin/main.go")
	if err != nil {
		panic(err)
	}

	// (os.Openによって呼び出された)ファイルを閉じる
	defer file.Close()

	// 標準出力へ書き出す
	io.Copy(os.Stdout, file)
}
