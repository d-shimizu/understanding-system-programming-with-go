package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	// サーバへ接続する
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	// リクエストを作成する
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888/",
		nil,
	)
	if err != nil {
		panic(err)
	}
	// リクエストを書き込む
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	// レスポンスを読み込む
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	// レスポンスをダンプする
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	// Responseヘッダが0ではないか、TransferEncodingが含まれているかをチェックする
	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("Wrong Transfer Encoding")
	}

	for {
		// サイズを取得する
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		// 16進数のサイズをパースする。サイズがゼロならクローズする。
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		// サイズ数分バッファを確保して読み込む。
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}
}
