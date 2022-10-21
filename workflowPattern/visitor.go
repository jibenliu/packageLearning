package workflowPattern

import (
	"fmt"
	"strconv"
)

// Element interface
type Element interface {
	Accept(v CarVisitor)
}

// Engine is ConcreteElement
type Engine struct{}

// Accept operation that takes a visitor as an argument
func (e Engine) Accept(v CarVisitor) {
	v.visitEngine(e)
}

// Wheel is ConcreteElement
type Wheel struct {
	Number int
}

// Accept operation
func (w Wheel) Accept(v CarVisitor) {
	v.visitWheel(w)
}

// Car is ConcreteElement
type Car struct {
	_items []Element
}

// Accept operation
func (c Car) Accept(v CarVisitor) {
	for _, e := range c._items {
		e.Accept(v)
	}
	v.visitCar(c)
}

// CarVisitor is Visitor
type CarVisitor interface {
	visitEngine(engine Engine)
	visitWheel(wheel Wheel)
	visitCar(car Car)
}

// TestCarVisitor is ConcreteVisitor
type TestCarVisitor struct{}

func (v TestCarVisitor) visitEngine(engine Engine) {
	fmt.Println("test engine")
}

func (v TestCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("test wheel #" +
		strconv.Itoa(wheel.Number))
}

func (v TestCarVisitor) visitCar(car Car) {
	fmt.Println("test car")
}

// RepairCarVisitor is ConcreteVisitor
type RepairCarVisitor struct{}

func (v RepairCarVisitor) visitEngine(engine Engine) {
	fmt.Println("repair engine")
}

func (v RepairCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("repair wheel #" +
		strconv.Itoa(wheel.Number))
}

func (v RepairCarVisitor) visitCar(car Car) {
	fmt.Println("repair car")
}

func VisitorDemo() {
	//Client
	car := Car{[]Element{
		Engine{},
		Wheel{1},
		Wheel{2},
		Wheel{3},
		Wheel{4},
	}}
	v1 := TestCarVisitor{}
	v2 := RepairCarVisitor{}

	car.Accept(v1)
	car.Accept(v2)

}
