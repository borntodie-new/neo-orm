package dialect

import (
	"fmt"
	"reflect"
	"time"
)

const (
	DriverName string = "sqlite3"
)

type sqlite3 struct{}

func init() {
	RegisterDialect(DriverName, &sqlite3{})
}

// DataTypeOf 将Go中的数据类型和SQL中的数据里的数据类型进行转换
func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// TableExistSQL 判断表是个否存在
func (s *sqlite3) TableExistSQL(tableName string) (string, []any) {
	args := []any{tableName}
	return "SELECT `name` FROM `sqlite_master` WHERE `type` = 'table' AND `name` = ?", args
}

var _ Dialect = (*sqlite3)(nil)
