package day02

import (
	"context"
	"github.com/borntodie-new/neo-orm/day02/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEngine(t *testing.T) {
	e, err := NewEngine(dialect.DriverName, "neo.db")
	assert.NoError(t, err)

	s := e.NewSession().Model(&TestModel{})
	ctx := context.Background()
	_ = s.DropTable(ctx)
	_ = s.CreateTable(ctx)
	if !s.HasTable(ctx) {
		t.Fatal("Failed to create table User")
	}

}

type TestModel struct {
	Id        int    `neoorm:"PRIMARY KEY"`
	FirstName string `neoorm:"AUTO INCREMENT"`
	Age       uint8
	LastName  string
}
