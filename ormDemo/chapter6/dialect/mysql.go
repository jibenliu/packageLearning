package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type Mysql struct {
}

func (m *Mysql) DataTypeOf(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return "enum(0,1)"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "int(10)"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "varchar(10)"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := v.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", v.Type().Name(), v.Kind()))
}

func (m *Mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT table_name FROM information_schema.TABLES WHERE table_name = ?", args
}

var _ Dialect = (*Mysql)(nil)

func init() {
	RegisterDialect("mysql", &Mysql{})
}
