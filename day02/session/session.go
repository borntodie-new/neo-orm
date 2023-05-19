package session

import (
	"context"
	"database/sql"
	"github.com/borntodie-new/neo-orm/day02/dialect"
	"github.com/borntodie-new/neo-orm/day02/log"
	"github.com/borntodie-new/neo-orm/day02/schema"
	"strings"
)

type Session struct {
	// db 维护住的数据库连接对象
	db *sql.DB
	//sql 用于拼接SQL语句
	sql strings.Builder
	// sqlVars SQL语句所需要的参数
	sqlVars []any
	// refTable 维护一个表模型
	refTable *schema.Schema
	// dialect 维护一个方言
	dialect dialect.Dialect
}

// New 创建一个Session对象，用于和数据库交互
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
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
func (s *Session) ExecContext(ctx context.Context) (res sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	res, err = s.DB().ExecContext(ctx, s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// QueryRowContext query row with context
func (s *Session) QueryRowContext(ctx context.Context) *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRowContext(ctx, s.sql.String(), s.sqlVars...)
}

// QueryRowsContext query rows with context
func (s *Session) QueryRowsContext(ctx context.Context) (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err = s.DB().QueryContext(ctx, s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
