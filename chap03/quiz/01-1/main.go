package main

import (
	"io"
	"os"
	"flag"
)

func main() {
    flag.Parse()
	old_file, err := os.Open(flag.Arg(0))
	//old_file, err := flag.Arg(0)
	if err != nil {
		panic(err)
	}

	defer old_file.Close()

	new_file, err := os.Create(flag.Arg(1))
	if err != nil {
		panic(err)
	}

	io.Copy(new_file, old_file)
}
