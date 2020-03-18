package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`Find images

Usage:
	%s [path to find]
`, os.Args[0])
		return
	}
	root := os.Args[1]

	// Walk関数は第一引数にディレクトリのPATHをstring形式で受け取り、第二引数でファイルまたはディレクトリが見つかった時に実行する関数を指定する
	// 第2引数の中でWalkFuncを呼び出す https://xn--go-hh0g6u.com/pkg/path/filepath/
	// 今回は FileInfo の情報を取得してその内容によって判定する
	// https://xn--go-hh0g6u.com/pkg/path/filepath/#Walk
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {  // ディレクトリの場合の処理
			if info.Name() == "_build" { // _build という名前のディレクトリの場合の処理
				return filepath.SkipDir
			}
			return nil
		}

		// 英数字の大文字を小文字にして拡張子を取得する
		ext := strings.ToLower(filepath.Ext(info.Name()))
		if imageSuffix[ext] {  // 拡張子が存在するときの条件式
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return nil
			}
			fmt.Printf("%s\n", rel)
		}
		return nil
	})

	if err != nil {
		fmt.Println(1, err)
	}
}
