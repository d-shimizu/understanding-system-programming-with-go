package main

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"io"
	"os"
)

func main() {
	var out io.Writer
	// 標準出力かどうかを確認する
	if isatty.IsTerminal(os.Stdout.Fd()) {     // isatty.IsTerminal = true (os.Stdout.Fd() -> 1)
		out = colorable.NewColorableStdout()
	} else {
		out = colorable.NewNonColorable(os.Stdout)
	}

	if isatty.IsTerminal(os.Stdin.Fd()) {     // isatty.IsTerminal(os.Stdin.Fd()) = true (os.Stdin.Fd() -> 0)
		// fmt.Println(isatty.IsTerminal(os.Stdin.Fd()))
		fmt.Fprintln(out, "stdin: terminal")
	} else {
		fmt.Println("stdin: pipe")
	}

	if isatty.IsTerminal(os.Stdout.Fd()) {    // isatty.IsTerminal(os.Stdout.Fd()) = true (os.Stdout.Fd() -> 1)
		fmt.Fprintln(out, "stdout: terminal")
	} else {
		fmt.Println("stdout: pipe")
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {    // isatty.IsTerminal(os.Stderr.Fd()) = true (os.Stdout.Fd() -> 2)
		fmt.Fprintln(out, "stderr: terminal")
	} else {
		fmt.Println("stderr: pipe")
	}
}
