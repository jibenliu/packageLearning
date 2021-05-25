package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...) //追加元素后如果超长就将原有数组长度翻倍
	fmt.Println("food:", food)
	fmt.Println("sliceLen:", len(food))//5
	fmt.Println("sliceCap:", cap(food))//6

	anyThingElse := []string{"oranges", "apples", "cat", "mouse", "dog", "desk", "keyboard", "shit"}//cap 8
	el := append(food, anyThingElse...)
	fmt.Println("bigSliceCap:", cap(el))//5+8=13 > 6*2 所以取13，否则的话翻倍也就是12
	fmt.Println("bigSliceLen:", len(el))

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
