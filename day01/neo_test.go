package day01

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestEngine(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	engine, _ := NewEngine("sqlite3", "neo.db")
	ctx := context.Background()
	defer engine.Close()
	s := engine.NewSession()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = s.Raw("DROP TABLE IF EXISTS User;").ExecContext(ctx)
			_, _ = s.Raw("CREATE TABLE User(Name text);").ExecContext(ctx)
			_, _ = s.Raw("CREATE TABLE User(Name text);").ExecContext(ctx)
			result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").ExecContext(ctx)
			count, _ := result.RowsAffected()
			fmt.Printf("Exec success, %d affected\n", count)
		})
	}

}
