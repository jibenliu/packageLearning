package main

import "time"

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name    string    `validate:"ne=admin"`
	Age     int       `validate:"gte=18"`
	Sex     string    `validate:"oneof=male female"`
	RegTime time.Time `validate:"lte"` //默认和当前时间对比
}

func main() {
	validate := validator.New()

	u1 := User{Name: "dj", Age: 18, Sex: "male", RegTime: time.Now().UTC()}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{Name: "admin", Age: 15, Sex: "none", RegTime: time.Now().UTC().Add(1 * time.Hour)}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
