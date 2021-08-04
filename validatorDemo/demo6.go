package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

//-：跳过该字段，不检验；
//|：使用多个约束，只需要满足其中一个，例如rgb|rgba；
//required：字段必须设置，不能为默认值；
//omitempty：如果字段未设置，则忽略它。

type User struct {
	Name  string `validate:"min=2"`
	Age   int    `validate:"min=18"`
	Email string `validate:"email"`
}

func main() {
	validate := validator.New()

	u1 := User{
		Name:  "dj",
		Age:   18,
		Email: "dj@example.com",
	}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{
		Name:  "dj",
		Age:   18,
		Email: "djexample.com",
	}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
