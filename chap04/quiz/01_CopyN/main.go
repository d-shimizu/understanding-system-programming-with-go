package main

import (
	"io"
	"os"
	"strings"
)

func copyn(sink io.Writer, src io.Reader, length int) {
	// 指定したバイト数を読み込む
	limitedSrc := io.LimitReader(src, int64(length))
	// sink(標準出力)にコピー
	io.Copy(sink, limitedSrc)
}

func main() {
	// 文字列を読み込む
	reader := strings.NewReader("Sample text 30.")
	copyn(os.Stdout, reader, 3)
}
