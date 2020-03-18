package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	// Unix Domain Socketへ接続する
	conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
	if err != nil {
		panic(err)
	}

	// localhost へリクエストを生成する
	request, err := http.NewRequest("get", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)

	// レスポンスを読み込む
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	// レスポンスをダンプする
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
