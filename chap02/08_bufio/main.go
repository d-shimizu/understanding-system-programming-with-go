package main

import (
	"bufio"
	"os"
)

func main() {
	// バッファを作成する
	buffer := bufio.NewWriter(os.Stdout)
	// バッファへ文字列を書き込む
	buffer.WriteString("bufio.Writer ")
	// io.Writerへのbufferを書き出す
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
