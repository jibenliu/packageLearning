package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var spinChars = `|/-\`

type Spinner struct {
	message string
	i       int
}

func NewSpinner(message string) *Spinner {
	return &Spinner{message: message}
}

func (s *Spinner) Tick() {
	fmt.Printf("%s %c \r", s.message, spinChars[s.i])
	s.i = (s.i + 1) % len(spinChars)
}

func isTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func main() {
	flag.Parse()
	s := NewSpinner("working...")
	for i := 0; i < 100; i++ {
		if isTTY() {
			s.Tick()
		}
		time.Sleep(100 * time.Millisecond)
	}

}
