package main

import "fmt"

func main() {
	var a = 10

	var p *int = &a

	fmt.Print("a =", a, "\n")
	fmt.Print("p =", p, "\n")
	fmt.Print("&p =", &p, "\n")

	a = 100
	fmt.Print(a, "\n")
	fmt.Print(p, "\n")
	fmt.Print(&p, "\n")

	*p = 200
	fmt.Print(a, "\n")
	fmt.Print(p, "\n")
	fmt.Print(&p, "\n")

}
