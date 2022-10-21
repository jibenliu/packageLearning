package workflowPattern

import (
	"fmt"
)

// Employee is Product
type Employee interface {
	Test()
}

// Booker is ConcreteProduct
type Booker struct{}

// Test is Employee method
func (e Booker) Test() {
	fmt.Println("Booker")
}

// Manager is ConcreteProduct
type Manager struct{}

// Test is Manager method
func (e Manager) Test() {
	fmt.Println("Manager")
}

// BookerCreator is ConcreteCreator
type BookerCreator struct{}

// CreateEmployee creates an Booker
func (c BookerCreator) CreateEmployee() Employee {
	return Booker{}
}

// ManagerCreator is ConcreteCreator
type ManagerCreator struct{}

// CreateEmployee creates an Manager
func (c ManagerCreator) CreateEmployee() Employee {
	return Manager{}
}

func FactoryMethodDemo() {
	//Client
	booker := BookerCreator{}.CreateEmployee()
	booker.Test()
	//printed: Booker

	manager := ManagerCreator{}.CreateEmployee()
	manager.Test()
	//printed: Manager
}
