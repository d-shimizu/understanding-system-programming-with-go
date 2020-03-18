package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	// スライスを作成
    sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}

	current := 0
	var conn net.Conn = nil

	// リトライ用にループで全体を囲う
	for {
		var err error
		// まだコネクションを貼っていない / エラーでリトライするときはDialから始まる
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// POSTで文字列を送るリクエストを作成する
		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// サーバから読み込む。タイムアウトはここでエラーになるのでリトライする。
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// 結果を表示する
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		current++

		if current == len(sendMessages) {
			break
		}

	}
}
