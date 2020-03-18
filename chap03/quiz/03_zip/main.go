package main

import (
	"archive/zip"
	"os"
	"io"
	"strings"
)

func main() {
	// .zipのファイルを作成する
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// zipファイル書き込み用の構造体を作成する
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// 圧縮対象ファイルを生成する
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("1つ目のテキストファイルです。"))

	// 圧縮対象ファイルを生成する
	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つ目のテキストファイルです。"))
}
