package seedwork

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type FakeDbSession struct {
	expectedSql    string
	expectedParams []any
	t              *testing.T
	rows           *FakeRows
}

func (s FakeDbSession) Exec(query string, args ...any) (infrastructure.Result, error) {
	assert.Equal(s.t, s.expectedSql, query)
	assert.Equal(s.t, s.expectedParams, args)
	return &DeferredResult{}, nil
}

func (s *FakeDbSession) Query(query string, args ...any) (infrastructure.Rows, error) {
	assert.Equal(s.t, s.expectedSql, query)
	assert.Equal(s.t, s.expectedParams, args)
	return s.rows, nil
}

func NewFakeRows(rows ...[]any) *FakeRows {
	return &FakeRows{
		rows, 0, false,
	}
}

type FakeRows struct {
	rows   [][]any
	idx    int
	Closed bool
}

func (r *FakeRows) Close() error {
	r.Closed = true
	return nil
}

func (r FakeRows) Err() error {
	return nil
}

func (r *FakeRows) Next() bool {
	r.idx++
	return len(r.rows) < r.idx
}

func (r FakeRows) Scan(dest ...any) error {
	for i, d := range dest {
		dt, ok := d.(sql.Scanner)
		if !ok {
			return errors.New("value should implement sql.Scanner interface")
		}
		err := dt.Scan(r.rows[r.idx][i])
		if err != nil {
			return err
		}
	}
	return nil
}

func TestMultiInsertQuery(t *testing.T) {
	cases := []struct {
		sql            string
		params         [][]any
		expectedSql    string
		expectedParams []any
	}{
		{"($1, $2, $3)",
			[][]any{[]any{1, 2, 3}, []any{4, 5, 6}},
			"($1, $2, $3), ($4, $5, $6)",
			[]any{1, 2, 3, 4, 5, 6},
		},
		{"($1, 'someone''s $2', $2)",
			[][]any{[]any{1, 2}, []any{3, 4}},
			"($1, 'someone''s $2', $2), ($3, 'someone''s $2', $4)",
			[]any{1, 2, 3, 4},
		},
		{"($1, '', $2)",
			[][]any{[]any{1, 2}, []any{3, 4}},
			"($1, '', $2), ($3, '', $4)",
			[]any{1, 2, 3, 4},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			sqlTemplate := `
                INSERT INTO tbname (
                    f1, f2, f3
                ) VALUES %s
                ON CONFLICT DO NOTHING`
			q := NewMultiInsertQuery()
			for _, v := range c.params {
				_, err := q.Exec(fmt.Sprintf(sqlTemplate, c.sql), v...)
				assert.Nil(t, err)
			}
			s := &FakeDbSession{
				fmt.Sprintf(sqlTemplate, c.expectedSql),
				c.expectedParams,
				t,
				NewFakeRows(),
			}
			_ = s
			_, err := q.Evaluate(s)
			assert.Nil(t, err)
		})
	}
}

func TestAutoincrementMultiInsertQuery(t *testing.T) {
	cases := []struct {
		sql            string
		params         [][]any
		expectedSql    string
		expectedParams []any
	}{
		{"($1, $2, $3)",
			[][]any{[]any{1, 2, 3}, []any{4, 5, 6}},
			"($1, $2, $3), ($4, $5, $6)",
			[]any{1, 2, 3, 4, 5, 6},
		},
		{"($1, 'someone''s $2', $2)",
			[][]any{[]any{1, 2}, []any{3, 4}},
			"($1, 'someone''s $2', $2), ($3, 'someone''s $2', $4)",
			[]any{1, 2, 3, 4},
		},
		{"($1, '', $2)",
			[][]any{[]any{1, 2}, []any{3, 4}},
			"($1, '', $2), ($3, '', $4)",
			[]any{1, 2, 3, 4},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			sqlTemplate := `
                INSERT INTO tbname (
                    f1, f2, f3
                ) VALUES %s
                RETURNING id`
			q := NewAutoincrementMultiInsertQuery()
			for _, v := range c.params {
				_, err := q.Exec(fmt.Sprintf(sqlTemplate, c.sql), v...)
				assert.Nil(t, err)
			}
			s := &FakeDbSession{
				fmt.Sprintf(sqlTemplate, c.expectedSql),
				c.expectedParams,
				t,
				NewFakeRows([]any{1}, []any{2}, []any{3}),
			}
			_ = s
			_, err := q.Evaluate(s)
			assert.Nil(t, err)
		})
	}
}
