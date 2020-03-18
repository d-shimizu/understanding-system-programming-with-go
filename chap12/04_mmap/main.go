package main

import (
	"github.com/edsrzf/mmap-go"
	"os"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

func main() {
	// テストデータの書き込み
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	if err != nil{
		panic(err)
	}

	// メモリにマッピングする
	// mは[]byteのエイリアスなので添字アクセス可能
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// mmap.Mapで指定した内容をメモリに書き込む
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	// mmap.Unmap(): メモリ上に展開された内容を削除して閉じる
	// deferにより最後に実行される
	defer m.Unmap()

	// メモリ上のデータを修正して書き込む
	m[9] = 'X'
	m.Flush()

	// 読み込んでみる
	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)

	// defer m.Unmap()
}
