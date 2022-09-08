package reflect_accelerate

import (
	"fmt"
	"reflect"
	"testing"
)

func populateStructReflect(in interface{}) error {
	val := reflect.ValueOf(in)
	if val.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("you must pass in a pointer")
	}
	elmv := val.Elem()
	if elmv.Type().Kind() != reflect.Struct {
		return fmt.Errorf("you must pass in a pointer to a struct")
	}

	fval := elmv.FieldByName("B")
	fval.SetInt(42)

	return nil
}

func BenchmarkPopulateReflect(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructReflect(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}
