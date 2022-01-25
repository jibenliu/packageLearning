package skipList

import (
	"fmt"
	"testing"
)

func Test_SkipList(t *testing.T) {
	sl := New()

	sl.Insert(float64(100), "foo")

	e, ok := sl.Search(float64(100))
	fmt.Println(ok)
	fmt.Println(e.Value)
	e, ok = sl.Search(float64(200))
	fmt.Println(ok)
	fmt.Println(e)

	sl.Insert(20.5, "bar")
	sl.Insert(float64(50), "spam")
	sl.Insert(float64(20), 42)

	fmt.Println(sl.Len())
	e = sl.Delete(float64(50))
	fmt.Println(e.Value)
	fmt.Println(sl.Len())

	for e := sl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Output:
	// true
	// foo
	// false
	// <nil>
	// 4
	// spam
	// 3
	// 42
	// bar
	// foo
}
