package main

import (
	"os"
	"io"
	"net"
)

func main() {
	// dshimizu.jp:80へ接続してコネクションを確立する
	conn, err := net.Dial("tcp", "dshimizu.jp:443")
	if err != nil {
		panic(err)
	}

	// 確立したコネクションへメッセージを送る
	conn.Write([]byte("Get / HTTP/1.1\r\nHost: dshimizu.jp\r\n\r\n"))

	// io.Reader インターフェースを利用してサーバから返ってきたメッセージを画面に出力する
	io.Copy(os.Stdout, conn)
}
