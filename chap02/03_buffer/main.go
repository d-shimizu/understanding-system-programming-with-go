package main

import (
	"fmt"
	"bytes"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
