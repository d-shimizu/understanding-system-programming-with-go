package main

import (
	"fmt"
	"os"
)

func main() {
	// 現在のディレクトリの情報を取得する関数
	wd, _ := os.Getwd()
	fmt.Println(wd)
}
