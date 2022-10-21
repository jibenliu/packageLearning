package workflowPattern

import (
	"fmt"
)

type state interface {
	open(c *Connection)
	close(c *Connection)
}

// CloseState is ConcreteState
type CloseState struct{}

func (cs CloseState) open(c *Connection) {
	fmt.Println("open the connection")
	c.setState(OpenState{})
}

func (cs CloseState) close(c *Connection) {
	fmt.Println("connection is already closed")
}

// OpenState is ConcreteState
type OpenState struct{}

func (os OpenState) open(c *Connection) {
	fmt.Println("connection is already open")
}

func (os OpenState) close(c *Connection) {
	fmt.Println("close the connection")
	c.setState(CloseState{})
}

// Connection is Context
type Connection struct {
	_state state
}

// Open connection
func (c *Connection) Open() {
	c._state.open(c)
}

// Close connection
func (c *Connection) Close() {
	c._state.close(c)
}

func (c *Connection) setState(state state) {
	c._state = state
}

func StateDemo() {
	//Client
	con := Connection{CloseState{}}
	con.Open()
	//printed: open the connection
	con.Open()
	//printed: connection is already open
	con.Close()
	//printed: close the connection
	con.Close()
	//printed: connection is already closed
}
