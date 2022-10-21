package workflowPattern

import (
	"fmt"
)

// Span is Flyweight
type Span interface {
	PrintSpan(style string)
}

// Char is ConcreteFlyweight
type Char struct {
	C rune
}

// PrintSpan is Operation(extrinsicState)
func (c Char) PrintSpan(style string) {
	fmt.Println("<span style=\"" +
		style + "\">" + string(c.C) + "</span>")
}

// CharFactory is FlyweightFactory
type CharFactory struct {
	chars map[rune]Char
}

// GetChar is GetFlyweight(key)
func (cf *CharFactory) GetChar(c rune) Span {
	if value, exists := cf.chars[c]; exists {
		return value
	}
	char := Char{c}
	cf.chars[c] = char
	return char
}

func FlyWeightDemo() {
	//Client
	factory := CharFactory{map[rune]Char{}}
	charA := factory.GetChar('A')
	charA.PrintSpan("font-size: 12")

	charB := factory.GetChar('B')
	charB.PrintSpan("font-size: 12")

	charA1 := factory.GetChar('A')
	charA1.PrintSpan("font-size: 12")

	equal := charA == charA1
	//equal is true

	fmt.Println(equal)
}
