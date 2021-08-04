package main

import (
	"fmt"
	"math"
)

const STR string = "constant"

func main() {
	fmt.Println(STR)
	const NUMBER = 500000

	const D = 3e20 / NUMBER
	fmt.Println(D)

	fmt.Println(int64(D))

	fmt.Println(math.Sin(NUMBER))
}
