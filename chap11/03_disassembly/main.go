package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {  // コマンドライン引数を受け取る
		fmt.Printf("%s [exec file name]", os.Args[0])
		// 引数がなければ終了する
		os.Exit(1)
	}

	// PATH環境変数の内容を取得して分割する関数 filepath.SplitList()
	for _, path := range filepath.SplitList(os.Getenv("PATH")){
		// 各PATHと引数の値を合体させる
		execpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execpath) // execpathに格納されたフルパス情報の有無をチェックする
		if !os.IsNotExist(err) {    // 存在していないファイルにアクセスしようとした際に得たエラー情報を渡すとTrueを返してくれる関数 os.IsNotExist()
			fmt.Println(execpath)
			return
		}
	}
	os.Exit(1)
}
