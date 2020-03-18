package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// ディレクトリのパスとファイル名とを連結する関数 filepath.Join()
	// システムの一時ファイル置き場のディレクトリパスを返す関数 os.TempDir()関数
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}
