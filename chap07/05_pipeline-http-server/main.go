
package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 順番に従ってconnに書き出しをする(goroutineで実行される) 順序整理用のキュー(待つ用)
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()

	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		response := <-sessionResponse // chanB <- chanA
		response.Write(conn)
		close(sessionResponse)
	}
}

// セッション内のリクエストを処理する (別なチャンネル内のチャンネルからのレスポンスを処理してクライアントに返すデータを作成する)
func handleRequest(request *http.Request, resultReciver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	content := "Hello World\n"

	// レスポンスを書き込む
	// セッションを維持するために Keep-Alive しないといけない
	response := &http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		ContentLength: int64(len(content)),
		Body: ioutil.NopCloser(strings.NewReader(content)),
	}

	// 処理が終わったらチャンネルに書き込む
	// ブロックされていたwriteToConnの処理を続ける
	resultReciver <- response
}

// セッション1つを処理する
func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// セッション内のリクエストを順に処理するチャンネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// レスポンスを直列化してソケットに書き出す専用のゴルーチン
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)

	for {
		// レスポンスを受け取ってキューに入れる
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		// リクエストを読み込む
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse

		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
