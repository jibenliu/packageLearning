package reflect_accelerate

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 优化二：利用字段偏移量
var unsafeCache = make(map[reflect.Type]uintptr)

type intface struct {
	typ   unsafe.Pointer
	value unsafe.Pointer
}

func populateStructUnsafe(in interface{}) error {
	typ := reflect.TypeOf(in)

	offset, ok := unsafeCache[typ]
	if !ok {
		if typ.Kind() != reflect.Ptr {
			return fmt.Errorf("you must pass in a pointer")
		}
		if typ.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("you must pass in a pointer to a struct")
		}
		f, ok := typ.Elem().FieldByName("B")
		if !ok {
			return fmt.Errorf("struct does not have field B")
		}
		if f.Type.Kind() != reflect.Int {
			return fmt.Errorf("field B should be an int")
		}
		offset = f.Offset
		unsafeCache[typ] = offset
	}

	structPtr := (*intface)(unsafe.Pointer(&in)).value
	*(*int)(unsafe.Pointer(uintptr(structPtr) + offset)) = 42

	return nil
}
