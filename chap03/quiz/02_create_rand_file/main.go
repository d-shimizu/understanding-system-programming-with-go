package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.CopyN(file, rand.Reader, 1024)
}
