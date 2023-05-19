package schema

import (
	"github.com/borntodie-new/neo-orm/day02/dialect"
	"reflect"
)

// Field 字段元信息
// 结构体上的每个字段的信息
type Field struct {
	// Name 字段名，Go中的名字
	Name string
	// Type 字段类型，Go中的类型
	Type string
	// Tag 字段的标签
	Tag string
}

// Schema 表模型——表的元数据信息
type Schema struct {
	// Model 当前结构体的类型
	Model any
	// Name 当前结构体的名字
	Name string
	// Fields 当前结构体中所有的字段
	Fields []*Field
	// FieldNames 当前结构体中的所有的字段的名字
	FieldNames []string
	// FieldMap 将Fields和FieldNames两者进行组合
	FieldMap map[string]*Field
}

// Parse 将结构体解析成表模型元数据
func Parse(dest any, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    modelType,
		Name:     modelType.Name(),
		FieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		fd := modelType.Field(i)
		if !fd.Anonymous && fd.IsExported() {
			field := &Field{
				Name: fd.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(fd.Type))),
			}
			if v, ok := fd.Tag.Lookup("neoorm"); ok {
				field.Tag = v
			}
			// 添加字段信息
			schema.Fields = append(schema.Fields, field)
			// 添加字段名
			schema.FieldNames = append(schema.FieldNames, fd.Name)
			// 添加map信息
			schema.FieldMap[fd.Name] = field
		}
	}
	return schema
}
