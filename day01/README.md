day01：SQL基础

### 一、初识SQLite

关于sqlite3，大家直接看菜鸟教程就好。教程手把手教授

[SQLite3](https://www.runoob.com/sqlite/sqlite-tutorial.html)

### 二、database/sql标准库的使用

关于`database/sql`内置模块的使用，大家看我在掘金上发布的文章即可，里面非常详细的介绍了。

[Golang的`database/sql`的使用](https://juejin.cn/post/7234450312726069306)

### 三、实现一个简单的log库

这个日志库有点点懵，主要是`SetLevel`方法，不太清楚里面的逻辑

`log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)`

这段代码其实不会很难理解

- 初始化一个log对象
- 第一个参数是日志输出位置
- 第二个参数是每条日志的前缀`[error]`左右两边的东西是一些样式
- 第三个参数是每条日志需要输出的标准信息。我们这里填的是记录日志的输出文件和产生时间

### 四、核心结构Session

Session这部分不难，它是和数据库驱动做交互的桥梁，向前对接ORM框架，向后对接database

```go
package session

import (
	"context"
	"database/sql"
	"github.com/borntodie-new/neo-orm/day01/log"
	"strings"
)

type Session struct {
	// db 维护住的数据库连接对象
	db *sql.DB
	//sql 用于拼接SQL语句
	sql strings.Builder
	// sqlVars SQL语句所需要的参数
	sqlVars []any
}

// New 创建一个Session对象，用于和数据库交互
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear 清除SQL和参数
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB 获取具体的数据库连接对象
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw 设置原生的SQL语句
func (s *Session) Raw(sql string, values ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteByte(' ')
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// ExecContext execute raw sql with args
func (s *Session) ExecContext(ctx context.Context) (sql.Result, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	res, err := s.DB().ExecContext(ctx, s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return res, err
}

// QueryRowContext query row with context
func (s *Session) QueryRowContext(ctx context.Context) *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRowContext(ctx, s.sql.String(), s.sqlVars...)
}

// QueryRowsContext query rows with context
func (s *Session) QueryRowsContext(ctx context.Context) (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err := s.DB().QueryContext(ctx, s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}
```

### 五、核心结构Engine

Engine这部分也不难，它是整个ORM框架的引擎，框架就是由它盘活的

不过这里的设计有点和咱们的思想不太一样：为什么不在Engine内部维护一个Session对象

我们都知道，像Engine和Session这类结构，在ORM框架中都是唯一的，也就是说一般不会发生变动的，
所以在Engine中维护一个唯一的Session是非常合理的。不太清楚这里兔兔为什么这样设计

```go
package day01

import (
	"database/sql"
	"github.com/borntodie-new/neo-orm/day01/log"
	"github.com/borntodie-new/neo-orm/day01/session"
)

type Engine struct {
	db *sql.DB
}

// NewEngine 初始化一个Engine对象，盘活这个框架
func NewEngine(driver string, driverSourceName string) (*Engine, error) {
	db, err := sql.Open(driver, driverSourceName)
	if err != nil {
		log.Error(err)
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
	}
	e := &Engine{db: db}
	log.Info("Connect database success")
	return e, nil
}

// Close 关闭db连接对象
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Fail to close the database")
	}
	log.Error("success to close the database")
}

// NewSession 创建一个Session对象
// 个人感觉Session对象可以在Engine中唯一维护一个就好
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
```

### 六、测试
测试就没什么东西了，直接看兔兔的文章就好