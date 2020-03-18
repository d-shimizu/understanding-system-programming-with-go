package main

import (
	"context"
	"fmt"
	"github.com/lestrrat/go-server-starter/listener"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// シグナル初期化
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// Server:Starterからもらったソケットを確認する
	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	// ウェブサーバをgoroutineで起動する
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Server pid %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	// SIGTERMを受け取ったら終了させる
	<-signals
	server.Shutdown(context.Background())
}
