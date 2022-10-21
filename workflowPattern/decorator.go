package workflowPattern

import (
	"fmt"
)

// DShape is Component
type DShape interface {
	ShowInfo()
}

// DSquare is ConcreteComponent
type DSquare struct{}

// ShowInfo is Operation()
func (s DSquare) ShowInfo() {
	fmt.Print("DSquare")
}

// DShapeDecorator is Decorator
type DShapeDecorator struct {
	DShape DShape
}

// ShowInfo is Operation()
func (sd DShapeDecorator) ShowInfo() {
	sd.DShape.ShowInfo()
}

// ColorDShape is ConcreteDecorator
type ColorDShape struct {
	DShapeDecorator
	color string
}

// ShowInfo is Operation()
func (cs ColorDShape) ShowInfo() {
	fmt.Print(cs.color + " ")
	cs.DShape.ShowInfo()
}

// ShadowDShape is ConcreteDecorator
type ShadowDShape struct {
	DShapeDecorator
}

// ShowInfo is Operation()
func (ss ShadowDShape) ShowInfo() {
	ss.DShape.ShowInfo()
	fmt.Print(" with shadow")
}

func DecoratorDemo() {
	//Client
	DSquare := DSquare{}
	DSquare.ShowInfo()
	//printed: DSquare
	fmt.Println()

	colorDShape := ColorDShape{
		DShapeDecorator{DSquare}, "red"}
	colorDShape.ShowInfo()
	//printed: red DSquare
	fmt.Println()

	shadowDShape := ShadowDShape{
		DShapeDecorator{colorDShape}}
	shadowDShape.ShowInfo()
	//printed: red DSquare with shadow
}
