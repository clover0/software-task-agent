package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Runner binary for %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
