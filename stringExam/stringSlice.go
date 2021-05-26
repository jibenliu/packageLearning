package main

import "fmt"

func main()  {
	s := "Hello World!"
	fmt.Println(s[0])         // 72
	fmt.Println(string(s[0])) // H

	fmt.Println(s[:5]) //Hello
	a := []rune(s)
	fmt.Println(a[0]) //72
	return
}
