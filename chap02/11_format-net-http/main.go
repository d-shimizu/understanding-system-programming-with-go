package main

import (
	"os"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://dshimizu.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
