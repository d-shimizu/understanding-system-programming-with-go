package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// 実行コマンドを変数定義
	count := exec.Command("./count")
	// 標準出力につながるパイプを取得する
	stdout, _ := count.StdoutPipe()

	go func() {
		// bufio.NewScanner ファイルまたは標準入力から1行ずつ読込する
		// https://qiita.com/mztnnrt/items/7b4982dc0e3b8fbc3e5f
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}
}
