package workflowPattern

import (
	"fmt"
)

// Command type
type Command interface {
	Execute()
}

// BankClient is Invoker
type BankClient struct {
	putCommand Command
	getCommand Command
}

// PutMoney runs the putCommand
func (bc BankClient) PutMoney() {
	bc.putCommand.Execute()
}

// GetMoney runs the getCommand
func (bc BankClient) GetMoney() {
	bc.getCommand.Execute()
}

// Bank is Receiver
type Bank struct{}

func (b Bank) giveMoney() {
	fmt.Println("money to the client")
}

func (b Bank) receiveMoney() {
	fmt.Println("money from the client")
}

// PutCommand is ConcreteCommand
type PutCommand struct {
	bank Bank
}

// Execute command
func (pc PutCommand) Execute() {
	pc.bank.receiveMoney()
}

// GetCommand is ConcreteCommand
type GetCommand struct {
	bank Bank
}

// Execute command
func (gc GetCommand) Execute() {
	gc.bank.giveMoney()
}

func Demo() {
	//Client
	bank := Bank{}
	cPut := PutCommand{bank}
	cGet := GetCommand{bank}
	client := BankClient{cPut, cGet}
	client.GetMoney()
	//printed: money to the client
	client.PutMoney()
	//printed: money from the client
}
