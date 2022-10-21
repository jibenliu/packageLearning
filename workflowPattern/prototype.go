package workflowPattern

import "fmt"

// PShape is Prototype
type PShape interface {
	Clone() PShape
}

// Square is ConcretePrototype
type Square struct {
	LineCount int
}

// Clone creates a copy of the square
func (s Square) Clone() PShape {
	return Square{s.LineCount}
}

//Client

// ShapeMaker contains a Shape
type ShapeMaker struct {
	Shape PShape
}

// MakeShape creates a copy of the Shape
func (sm ShapeMaker) MakeShape() PShape {
	return sm.Shape.Clone()
}

func PrototypeDemo() {
	square := Square{4}
	maker := ShapeMaker{square}

	square1 := maker.MakeShape()
	square2 := maker.MakeShape()

	fmt.Println(square1)
	fmt.Println(square2)
}
