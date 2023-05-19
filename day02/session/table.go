package session

import (
	"context"
	"fmt"
	"github.com/borntodie-new/neo-orm/day02/log"
	"github.com/borntodie-new/neo-orm/day02/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value any) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}
func (s *Session) RefTable() *schema.Schema {
	if s.refTable != nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable(ctx context.Context) error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ", ")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).ExecContext(ctx)
	return err
}

func (s *Session) DropTable(ctx context.Context) error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.refTable.Name)).ExecContext(ctx)
	return err
}

func (s *Session) HasTable(ctx context.Context) bool {
	sql, values := s.dialect.TableExistSQL(s.refTable.Name)
	raw := s.Raw(sql, values...).QueryRowContext(ctx)
	var temp string
	_ = raw.Scan(&temp)
	return temp == s.RefTable().Name
}
