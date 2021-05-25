package main

import "fmt"

type Callable func()

type handle = int

func callableTest(val interface{}) {
	switch v := val.(type) {
	case func():
		v()
	default:
		fmt.Println("wrong type")
	}
}
func main() {
	myCallable := func() {
		fmt.Println("callable")
	}
	var h handle = 2
	callableTest(myCallable)
	callableTest(h)

	functions := make([]func(), 3)
	for i := 0; i < 3; i++ { //迭代时go仅获取了变量的指针
		functions[i] = func() {
			fmt.Println(fmt.Sprintf("iterator value: %d", i))
		}
	}

	functions[0]()
	functions[1]()
	functions[2]()

	functionsNew := make([]func(), 3)
	for i := 0; i < 3; i++ {
		functionsNew[i] = func(y int) func() {
			return func() {
				fmt.Println(fmt.Sprintf("iterator value: %d", y))
			}
		}(i)
	}

	functionsNew[0]()
	functionsNew[1]()
	functionsNew[2]()


	functionsNew1 := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // Trick mit neuer Variable  循环的每次迭代中都会创建一个新的变量
		functionsNew1[i] = func() {
			fmt.Println(fmt.Sprintf("iterator value: %d", i))
		}
	}

	functionsNew1[0]()
	functionsNew1[1]()
	functionsNew1[2]()
}
