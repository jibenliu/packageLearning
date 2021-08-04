package cls1

import (
	"fmt"
	"learnGo/factoryMode/base"
)

type Class1 struct {
}

func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
