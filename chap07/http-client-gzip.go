package main

import (
	"fmt"
	"compress/gzip"
	"buffio"
	"net"
	"net/http"
	"net/http/httputil"
	"string"
)

func main() {
	sendMessage := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}

	current := 0
	var conn net.Conn = nil

	for {
		var err error

		// まだコネクションを貼っていない / エラーでリトライ時はDialから行う
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Access: %d\n", current)
		}

		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
							"POST",
							"http://localhost:8888",
							strings.NewReader(sendMessage[current])
						)
		if err != nil {
			panic(err)
		}
		request.Handler.Set("Accept-Encoding", "gzip")

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// サーバから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(bufio.NewReader(conn), request)

		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// 結果を表示
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		fmt.Println(string(response.Body))

		defer response.Body.Close()
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

		// 全部送信完了していれば終了
		current++
		if current == len(sendMessage){
			break
		}

	}
}
