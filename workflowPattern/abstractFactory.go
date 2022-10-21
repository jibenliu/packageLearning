package workflowPattern

import (
	"fmt"
)

// ProductA is abstract product A
type ProductA interface {
	TestA()
}

// ProductB is abstract product B
type ProductB interface {
	TestB()
}

// Factory is abstract factory
type Factory interface {
	CreateA() ProductA
	CreateB() ProductB
}

// ProductA1 is concrete product A1
type ProductA1 struct{}

// TestA is implementation of the ProductA interface method
func (p ProductA1) TestA() {
	fmt.Println("test A1")
}

// ProductB1 is concrete product B1
type ProductB1 struct{}

// TestB is implementation of the ProductB interface method
func (p ProductB1) TestB() {
	fmt.Println("test B1")
}

// ProductA2 is concrete product A2
type ProductA2 struct{}

// TestA is implementation of the ProductA interface method
func (p ProductA2) TestA() {
	fmt.Println("test A2")
}

// ProductB2 is concrete product B2
type ProductB2 struct{}

// TestB is implementation of the ProductB interface method
func (p ProductB2) TestB() {
	fmt.Println("test B2")
}

// Factory1 is concrete factory 1
type Factory1 struct{}

// CreateA is implementation of the Factory interface method
func (f Factory1) CreateA() ProductA {
	return ProductA1{}
}

// CreateB is implementation of the Factory interface method
func (f Factory1) CreateB() ProductB {
	return ProductB1{}
}

// Factory2 is concrete factory 2
type Factory2 struct{}

// CreateA is implementation of the Factory interface method
func (f Factory2) CreateA() ProductA {
	return ProductA2{}
}

// CreateB is implementation of the Factory interface method
func (f Factory2) CreateB() ProductB {
	return ProductB2{}
}

// TestFactory creates and tests factories
func TestFactory(factory Factory) {
	productA := factory.CreateA()
	productB := factory.CreateB()
	productA.TestA()
	productB.TestB()
}

func AbstractFactoryDemo() {
	//client code:
	TestFactory(Factory1{})
	//printed: test A1
	//         test B1
	TestFactory(Factory2{})
	//printed: test A2
	//         test B2
}
