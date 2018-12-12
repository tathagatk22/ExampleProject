package main

import (
	"fmt"
	"golang.org/x/tools/cmd/guru/testdata/src/alias"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("\\opt\\app-root\\src", "c"))
	fmt.Println(filepath.Join("a/b"))
}
