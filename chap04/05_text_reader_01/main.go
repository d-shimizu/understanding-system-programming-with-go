package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目
`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		// \nのところで文字列を分割する
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", line)
		break
	}
}
