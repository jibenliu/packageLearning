package main

import (
	"fmt"
	"github.com/jibenliu/govaluateplus"
)

type Foo struct {
	Length int
}

func (f Foo) Echo(str string) int {
	fmt.Println(str)
	return 1
}

func main() {
	var foo = Foo{Length: 200}
	expression, _ := govaluateplus.NewEvaluableExpression("foo.Echo('hello world')")
	parameters := make(map[string]interface{}, 1)
	parameters["foo"] = foo
	result, err := expression.Evaluate(parameters)
	fmt.Println(result, err)

	expression, _ = govaluateplus.NewEvaluableExpression("foo.Length > 9000")
	result, err = expression.Evaluate(parameters)
	fmt.Println(result, err)
	functions := map[string]govaluateplus.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
	}

	expString := "strlen('someReallyLongInputString') <= 16"
	expression, _ = govaluateplus.NewEvaluableExpressionWithFunctions(expString, functions)

	result, _ = expression.Evaluate(nil)
	fmt.Println(result)
}
