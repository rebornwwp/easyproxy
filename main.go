package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("hello")
	fmt.Println()
	fmt.Println(filepath.Join("hello", "world"))
}
