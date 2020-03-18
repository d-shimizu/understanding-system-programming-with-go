package main

import (
    "bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// localhost:8888へ接続する。
    conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println(conn) // ポインタの情報？

	// リクエストを作成する
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
    request.Write(conn)

	// レスポンスを読み込む
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	// レスポンスをダンプする
	dump, err := httputil.DumpResponse(response, true)
    if err != nil {
	    panic(err)
	}
	fmt.Println(string(dump))
}
