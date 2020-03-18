package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// パスからファイル名とその親ディレクトリに分割する関数: filepath.Split()
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}
