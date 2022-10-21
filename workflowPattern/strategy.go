package workflowPattern

import "fmt"

// Strategy with integer operation
type Strategy interface {
	DoOperation(a int, b int) int
}

// AddStrategy is ConcreteStrategy
type AddStrategy struct{}

// DoOperation is integer addition function
func (s AddStrategy) DoOperation(a int, b int) int {
	return a + b
}

// SubstractStrategy is ConcreteStrategy
type SubstractStrategy struct{}

// DoOperation is integer subtraction function
func (s SubstractStrategy) DoOperation(a int, b int) int {
	return a - b
}

// Calc is Context
type Calc struct {
	_strategy Strategy
}

// Execute current strategy
func (c Calc) Execute(a int, b int) int {
	if c._strategy == nil {
		return 0
	}

	return c._strategy.DoOperation(a, b)
}

// SetStrategy changes the current strategy
func (c *Calc) SetStrategy(strategy Strategy) {
	c._strategy = strategy
}

func StrategyDemo() {
	//Client
	calc := Calc{}
	result1 := calc.Execute(5, 3)
	//result1 is 0

	calc.SetStrategy(AddStrategy{})
	result2 := calc.Execute(5, 3)
	//result2 is 8

	calc.SetStrategy(SubstractStrategy{})
	result3 := calc.Execute(5, 3)
	//result3 is 2

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
