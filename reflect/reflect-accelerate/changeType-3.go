package reflect_accelerate

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 优化三：更改缓存 key 类型
var unsafeCache2 = make(map[uintptr]uintptr)

func populateStructUnsafe2(in interface{}) error {
	inf := (*intface)(unsafe.Pointer(&in))

	offset, ok := unsafeCache2[uintptr(inf.typ)]
	if !ok {
		typ := reflect.TypeOf(in)
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
		unsafeCache2[uintptr(inf.typ)] = offset
	}

	*(*int)(unsafe.Pointer(uintptr(inf.value) + offset)) = 42

	return nil
}
