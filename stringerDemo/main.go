package main

import (
	"fmt"
	"stringerDemo/mycodes"
)

//go:generate stringer -linecomment -type ErrCode ./mycodes
func main() {
	fmt.Println(mycodes.BadRequest)
}
