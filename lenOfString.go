package main

import (
	"fmt"
)

func main() {
	theme := "狙击 start"

	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}

	for _, s := range theme {
		fmt.Printf("Unicode:%c %d\n", s, s)
	}
}
