package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

//contains=：包含参数子串，例如contains=email；
//containsany：包含参数中任意的 UNICODE 字符，例如containsany=abcd；
//containsrune：包含参数表示的 rune 字符，例如containsrune=☻；
//excludes：不包含参数子串，例如excludes=email；
//excludesall：不包含参数中任意的 UNICODE 字符，例如excludesall=abcd；
//excludesrune：不包含参数表示的 rune 字符，excludesrune=☻；
//startswith：以参数子串为前缀，例如startswith=hello；
//endswith：以参数子串为后缀，例如endswith=bye。

type User struct {
	Name     string `validate:"containsrune=☻"`
	Age      int    `validate:"min=18"`
	PlatCode string `validate:"startswith=prefix_"`
}

func main() {
	validate := validator.New()

	u1 := User{"d☻j", 18, "prefix_mm"}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{"dj", 18, "test"}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
