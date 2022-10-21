package workflowPattern

import (
	"fmt"
)

// Complex parts
type kettle struct{}

func (k kettle) TurnOff() {
	fmt.Println("Kettle turn off")
}

type toaster struct{}

func (t toaster) TurnOff() {
	fmt.Println("Toaster turn off")
}

type refrigerator struct{}

func (r refrigerator) TurnOff() {
	fmt.Println("Refrigerator turn off")
}

// Facade
type kitchen struct {
	kettle       kettle
	toaster      toaster
	refrigerator refrigerator
}

func (k kitchen) Off() {
	k.kettle.TurnOff()
	k.toaster.TurnOff()
	k.refrigerator.TurnOff()
}

func FacadeDemo() {
	//Client
	kitchen := kitchen{
		kettle{},
		toaster{},
		refrigerator{},
	}

	kitchen.Off()
}
