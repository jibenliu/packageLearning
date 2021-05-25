package main

import (
	"fmt"
)

type Strategy interface {
	StrategyMethod()
}
type ConcreteStrategyA struct {
}

func (s ConcreteStrategyA) StrategyMethod() {
	fmt.Println("concrete strategy A's method is visited.")
}

type ConcreteStrategyB struct {
}

func (s ConcreteStrategyB) StrategyMethod() {
	fmt.Println("concrete strategy B's method is visited.")
}

type Context struct {
	strategy Strategy
}

func (cx *Context) SetStrategy(strategy Strategy) {
	cx.strategy = strategy
}
func (cx *Context) GetStrategy() Strategy {
	return cx.strategy
}
func (cx *Context) StrategyMethod() {
	cx.strategy.StrategyMethod()
}

func main() {
	context := new(Context)
	var s Strategy = new(ConcreteStrategyA)
	context.SetStrategy(s)
	context.StrategyMethod()
	fmt.Println("================================")
	s = new(ConcreteStrategyB)
	context.SetStrategy(s)
	context.StrategyMethod()
}
