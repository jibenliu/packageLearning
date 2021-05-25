package main

import "fmt"

type State int

const (
	// Unknown State = iota
	Running State = iota + 1 //避免默认值被判定未Running ,或者采取上一行的方法
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

func main() {
	state := Running

	fmt.Println("state ", state)
}
