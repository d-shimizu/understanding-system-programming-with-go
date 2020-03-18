package main

import (
	"fmt"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"io/ioutil"
	"strings"
)

func main() {
	// observableを作成する
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	// イベントを受け取るobserverを作成する
	// https://godoc.org/github.com/ReactiveX/RxGo/observer#pkg-variables
	// Observer構造体へのアクセス
	// observerパッケージのObserver構造体へアクセスする
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Println("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done")
		},
	}

	// observableとobserverを接続する
	sub := source.Subscribe(watcher)

	// observableに値を投入する
	go func() {
		// main.goを読み込む
		content, err := ioutil.ReadFile("main.go")
		if err != nil {
			emitter <- err
		} else {
			// 改行で区切る
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()
	// 終了待ち
	<-sub
}
