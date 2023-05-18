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
