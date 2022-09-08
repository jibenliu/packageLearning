package reflect_accelerate

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type typeDescriptor uintptr

func describeType(in interface{}) (typeDescriptor, error) {
	typ := reflect.TypeOf(in)
	if typ.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("you must pass in a pointer")
	}
	if typ.Elem().Kind() != reflect.Struct {
		return 0, fmt.Errorf("you must pass in a pointer to a struct")
	}
	f, ok := typ.Elem().FieldByName("B")
	if !ok {
		return 0, fmt.Errorf("struct does not have field B")
	}
	if f.Type.Kind() != reflect.Int {
		return 0, fmt.Errorf("field B should be an int")
	}
	return typeDescriptor(f.Offset), nil
}

func populateStructUnsafe3(in interface{}, ti typeDescriptor) error {
	structPtr := (*intface)(unsafe.Pointer(&in)).value
	*(*int)(unsafe.Pointer(uintptr(structPtr) + uintptr(ti))) = 42
	return nil
}

func BenchmarkPopulateUnsafe3(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct

	descriptor, err := describeType((*SimpleStruct)(nil))
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe3(&m, descriptor); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}
