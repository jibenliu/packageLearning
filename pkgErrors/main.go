package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func DiffSub2(foo int, bar int) error {
	return errors.New("sub2 diff error")
}

func DiffSub1(foo int, bar int) error {
	if foo < 0 {
		return errors.New("diff error")
	}
	if err := DiffSub2(foo, bar); err != nil {
		return errors.Wrap(err, "sub1 error:")
	}
	return nil
}

func main() {
	err := DiffSub1(1, 2)
	fmt.Printf("%+v", err)
}
