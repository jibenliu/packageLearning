package main

import "reflect"

func IsInSlice(value int, sli []int) bool {
	for _, v := range sli {
		if value == v {
			return true
		}
	}
	return false
}

func upgrade(value interface{}, sli interface{}) bool {
	switch reflect.TypeOf(sli).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(sli)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
