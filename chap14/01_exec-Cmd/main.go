package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	state := cmd.ProcessState
	// 終了コードと状態を文字列で返す
	fmt.Printf("%s\n", state.String())
	// 子プロセスのプロセスID
	fmt.Printf("	Pid: %d\n", state.Pid())
	// カーネル内で消費された時間
	fmt.Printf("	System: %d\n", state.SystemTime())
	// ユーザランドで消費された時間
	fmt.Printf("	User: %d\n", state.UserTime())
}
