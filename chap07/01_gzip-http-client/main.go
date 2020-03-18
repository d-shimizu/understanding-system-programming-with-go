package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"io"
	"compress/gzip"
	"os"
	"strings"
)

func main() {
	// スライスを作成する
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}

	current := 0
	var conn net.Conn = nil

	for {
		var err error
		// コネクションが張られているかどうかのチェック
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成する 文字列はスライスの内容から取得する
		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		request.Header.Set("Accept-Encoding", "gzip")

		if err != nil {
			panic(err)
		}
		// リクエストを書き込む
		// _, err := request.Write(conn)
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// サーバから読み込む。タイムアウトした際はここでエラーになるためリトライする
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// 結果を表示する / 引数にfalseを指定することでResponse.Bodyを無視する
		dump, err := httputil.DumpResponse(response, false) // func DumpResponse(resp *http.Response, body bool) ([]byte, error)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()

		// gzip圧縮されている場合は展開して画面に出力する
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++
		// currentがスライスと同じ長さになった時終了
		if current == len(sendMessages) {
			break
		}
	}
}
