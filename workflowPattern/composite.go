package workflowPattern

import (
	"fmt"
)

// Graphic is Component
type Graphic interface {
	Draw()
}

// Сircle is Leaf
type Сircle struct{}

// Draw is Operation
func (c Сircle) Draw() {
	fmt.Println("Draw circle")
}

// CSquare is Leaf
type CSquare struct{}

// Draw is Operation
func (s CSquare) Draw() {
	fmt.Println("Draw CSquare")
}

// Image is Composite
type Image struct {
	graphics []Graphic
}

// Add Adds a Leaf to the Composite.
func (i *Image) Add(graphic Graphic) {
	i.graphics = append(i.graphics, graphic)
}

// Draw is Operation
func (i Image) Draw() {
	fmt.Println("Draw image")
	for _, g := range i.graphics {
		g.Draw()
	}
}

func CompositeDemo() {
	//Client
	image := Image{}
	image.Add(Сircle{})
	image.Add(CSquare{})
	picture := Image{}
	picture.Add(image)
	picture.Add(Image{})
	picture.Draw()
}
