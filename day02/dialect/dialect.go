package dialect

import "reflect"

// dialectsMap 维护一个全局的方言数据
var dialectsMap = map[string]Dialect{}

// Dialect 方言接口
// 在框架层面处理各个数据库产品之间的差异
type Dialect interface {
	// DataTypeOf 将Go中的数据类型转换成SQL中的数据类型
	DataTypeOf(typ reflect.Value) string
	// TableExistSQL 判断当前表是否以存在于数据库中
	TableExistSQL(tableName string) (string, []any)
}

// RegisterDialect 注册数据库
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取数据库
func GetDialect(name string) (Dialect, bool) {
	dialect, ok := dialectsMap[name]
	return dialect, ok
}
