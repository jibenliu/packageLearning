package main

import (
	"fmt"
	"strings"
)

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formatter",
		"tester name",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.Title,
	}

	StringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
