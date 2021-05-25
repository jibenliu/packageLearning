package main

import (
	"addBuildInfo/bininfo"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "show bin info")
	flag.Parse()
	if *v {
		_, _ = fmt.Fprint(os.Stderr, bininfo.StringifyMultiLine())
		os.Exit(1)
	}

	fmt.Println("my app running...")
	fmt.Println("bye...")
}
