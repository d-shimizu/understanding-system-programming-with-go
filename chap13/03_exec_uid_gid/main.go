package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	fmt.Printf("実行ユーザーID: %d\n",os.Geteuid())
	fmt.Printf("実行グループID: %d\n", os.Getegid())
}
