package main

import (
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	defer stdtty.Close()

	cmd.Stdin = stdpty
	cmd.Stdout = stdpty
	errpty, errtty, _ := pty.Open()

	defer errtty.Close()
	cmd.Stderr = errtty

	go func() {
		io.Copy(os.Stdout, stdpty)
	}()

	go func() {
		io.Copy(os.Stderr, errpty)
	}()

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
