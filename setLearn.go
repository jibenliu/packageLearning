package main

import "fmt"

//实现set集合，不能出现元素

var Exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

func New(items ...interface{}) *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = Exists
	}
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

type setTest struct{}

func main() {
	m1 := make(map[interface{}]struct{})
	m1["name"] = setTest{}
	m1["age"] = setTest{}
	m1["distinct"] = setTest{}
	m1["area"] = setTest{}
	s := Set{m: m1}

	m2 := make(map[interface{}]struct{})
	m2["distinct"] = setTest{}
	m2["area"] = setTest{}
	o := Set{m: m2}

	fmt.Println(s.Equal(&o))
	fmt.Println(o.IsSubset(&s))
}
