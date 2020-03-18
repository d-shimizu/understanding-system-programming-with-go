package main

import (
    "bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"io"
	"time"
)

func main() {
    // httpサーバを起動する
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
        panic(err)
	}
    fmt.Println("Server is runing at localhost:8888")

	// ループ
	for {
		// リクエストを受け付ける(まで待つことになる)
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// goroutine
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			for {
				// タイムアウト設定する
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				// リクエストを読み込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					// それ以外はエラーとする
					neterr, ok := err.(net.Error)  // ダウンキャスト(型アサーション)してタイムアウトの情報を取得する
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
			     }

				// リクエストを表示する
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
				    panic(err)
				}
				fmt.Println(string(dump))

				content := "Hello World\n"
				// レスポンスを書き込む
				// HTTP/1.1かつ、ContentLengthの設定が必要
				response := http.Response{
				    StatusCode: 200,
					ProtoMajor:   1,
					ProtoMinor:   1,
					ContentLength: int64(len(content)),
				}
				response.Write(conn)
			}
			conn.Close()
		}()
	}
}
