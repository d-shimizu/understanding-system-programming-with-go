package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}

	for _, task := range tasks {
		go func() {
			// goroutine が起動するときにはループが回り切って全部のタスクが最後のタスクになってしまう
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
