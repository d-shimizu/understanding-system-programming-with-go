package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprint(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
