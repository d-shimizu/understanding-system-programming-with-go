package main

import (
	"io"
	"os"
)

func main() {
	old_file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}

	defer old_file.Close()

	new_file, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(new_file, old_file)
}
