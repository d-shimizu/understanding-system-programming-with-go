package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(path string) chan string {
	// ファイルを読み込み、その結果を返すFutureを返す
	promise := make(chan string)
	go func() {
		// ファイルを読み込む
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			close(promise)
		} else {
			// 約束を果たした
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	// 文字列中の関数一覧を返すFutureを返す
	promise := make(chan []string)
	go func() {
		var result []string
		// future が解決するまで待って実行する
		// 受信チャネルからmain.goの内容を受け取り改行で分割する
		for _, line := range strings.Split(<-futureSource, "\n") {
			// funcで始まる文字列があるかどうかをチェックする
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		// 約束を果たした
		promise <- result
	}()
	return promise
}

func main() {
	// futureSource := readFile("future_promise.go")
	futureSource := readFile("main.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
