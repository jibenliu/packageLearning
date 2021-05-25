package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Match(`?`, "a"))
	fmt.Println(filepath.Match(`?`, "aa"))
	fmt.Println(filepath.Match(`a?`, "aa"))
	fmt.Println(filepath.Match(`a?`, "a"))
	fmt.Println(filepath.Match(`??`, "aa"))
	fmt.Println(filepath.Match(`*`, ""))
	fmt.Println(filepath.Match(`*`, "/a/b/c"))
	fmt.Println(filepath.Match(`//b`, "b"))
}
