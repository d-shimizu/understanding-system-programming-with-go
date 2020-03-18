package main

import (
	"io"
	"archive/zip"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment/; filename=sample.zip")

	// zipファイルを生成する
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// ファイルの数だけ書き込み
	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("1つ目のファイルのテキストです。"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("2つ目のファイルのテキストです。"))

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
