package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "blog.dshimizu.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: blog.dshimizu.jp\r\n\r\n"))

	// レスポンスのヘッダの内容を読み取る
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)

	// ヘッダーを表示する
	fmt.Println(res.Header)
	// fmt.Println(res.Body)

	// ボディーを表示する&ファイルをクローズする
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
