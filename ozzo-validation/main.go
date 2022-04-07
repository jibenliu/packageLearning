package main

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
)

func main() {
	data := "example"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		validation.Match(regexp.MustCompile("^[A-Z]{2}$")),
		validation.In("Female", "Male"),
		is.URL,   // is a valid URL
		is.Email, // is a valid URL
	)
	fmt.Println(err)
	// Output:
	// must be a valid URL

}

type Customer struct {
	Unit     int
	Quantity string
	Email    string
	Phone    string
	Name     string
	sth      string
	Emails   []string
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)),
		// Emails are optional, but if given must be valid.
		validation.Field(&c.Emails, validation.Each(is.Email)),
		validation.Field(&c.Unit, validation.When(c.Quantity != "", validation.Required).Else(validation.Nil)),
		validation.Field(&c.Phone, validation.When(c.Email == "", validation.Required.Error("Either phone or Email is required."))),
		validation.Field(&c.Email, validation.When(c.Phone == "", validation.Required.Error("Either phone or Email is required."))),
		validation.Field(&c.sth, NameRule...),
	)
}

func stringEquals(str string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != str {
			return errors.New("unexpected string")
		}
		return nil
	}
}

var NameRule = []validation.Rule{
	validation.Required,
	validation.Length(5, 20),
}

func validateCustomer() {
	c := Customer{
		Name: "Qiang Xue",
		Emails: []string{
			"valid@example.com",
			"invalid",
		},
	}

	err := c.Validate()
	fmt.Println(err)
}
