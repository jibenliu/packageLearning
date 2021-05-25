package main

import (
	"flag"
	"fmt"
	"os"
)

// go run demo1.go -a c
func main() {
	//str := os.Args[0]
	//fmt.Println(str)
	//fmt.Println(os.Args[1:])
	//
	//var cmdStr string
	//flag.StringVar(&cmdStr, "a", "default", "defaultStr")
	ptr := flag.String("a", "default", "defaultStr")
	flag.Parse()
	//fmt.Println(cmdStr)
	fmt.Println(*ptr)


	var name string
	for index, arg := range os.Args {
		if arg == "-a" {
			if index < len(os.Args)-1 {
				name = os.Args[index+1]
				break
			}
		}
	}
	fmt.Printf("a : %s\n", name)
}
