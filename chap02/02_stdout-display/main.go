package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("bytes.Buffer example\n"))
}
