package schema

import (
	"github.com/borntodie-new/neo-orm/day02/dialect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	d, ok := dialect.GetDialect(dialect.DriverName)
	assert.True(t, ok)
	testCases := []struct {
		name string
		dest any
		d    dialect.Dialect

		wantName string
	}{
		{
			name:     "test",
			dest:     &TestModel{},
			d:        d,
			wantName: "TestModel",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			schema := Parse(tc.dest, tc.d)
			assert.Equal(t, tc.wantName, schema.Name)
		})
	}
}

type TestModel struct {
	Id        int    `neoorm:"PRIMARY KEY"`
	FirstName string `neoorm:"AUTO INCREMENT"`
	Age       uint8
	LastName  string
}
