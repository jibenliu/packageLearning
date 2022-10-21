package workflowPattern

import "fmt"

// Expression is AbstractExpression
type Expression interface {
	Interpret(i int) bool
}

// DivExpression is TerminalExpression
type DivExpression struct {
	divider int
}

// Interpret func calculates expression
func (e DivExpression) Interpret(i int) bool {
	return i%e.divider == 0
}

// OrExpression is NonterminalExpression
type OrExpression struct {
	exp1 Expression
	exp2 Expression
}

// Interpret func calculates expression
func (e OrExpression) Interpret(i int) bool {
	return e.exp1.Interpret(i) || e.exp2.Interpret(i)
}

// AndExpression is NonterminalExpression
type AndExpression struct {
	exp1 Expression
	exp2 Expression
}

// Interpret func calculates expression
func (e AndExpression) Interpret(i int) bool {
	return e.exp1.Interpret(i) && e.exp2.Interpret(i)
}

func InterpreterDemo() {
	//Client
	divExp5 := DivExpression{5}
	divExp7 := DivExpression{7}
	orExp := OrExpression{
		divExp5, divExp7}
	andExp := AndExpression{
		divExp5, divExp7}

	//21 is divided by 7 or 5?
	result1 := orExp.Interpret(21)
	//result1 is true

	//21 is divided by 7 and 5?
	result2 := andExp.Interpret(21)
	//result2 is false

	//35 is divided by 7 and 5?
	result3 := andExp.Interpret(35)
	//result3 is true

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
