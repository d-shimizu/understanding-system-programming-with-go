package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 現在の階層を示す . や親ディレクトリの階層を示す .. を削除してきれいにする filepath.Clean() 関数
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// 絶対PATHを表す関数 filepath.Abs()
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	//  /usr/local/go/src/path/ilepath/path_unix.go

	// 相対PATHを表す関数 filepath.Rel()
	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)
	// path/filepath/path.go
}
