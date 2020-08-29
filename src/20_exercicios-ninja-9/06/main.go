package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operation System :: ", runtime.GOOS)
	fmt.Println("Operation System Arch ::", runtime.GOARCH)
}
