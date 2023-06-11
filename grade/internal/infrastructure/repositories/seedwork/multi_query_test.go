package seedwork

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type FakeDbSession struct {
	expectedSql    string
	expectedParams []any
	t              *testing.T
}

func (s FakeDbSession) Exec(query string, args ...any) (infrastructure.Result, error) {
	assert.Equal(s.t, s.expectedSql, query)
	assert.Equal(s.t, s.expectedParams, args)
	return &DeferredResult{}, nil
}

func (s FakeDbSession) Query(query string, args ...any) (infrastructure.Rows, error) {
	return &sql.Rows{}, nil
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
			s := FakeDbSession{
				fmt.Sprintf(sqlTemplate, c.expectedSql),
				c.expectedParams,
				t,
			}
			_ = s
			_, err := q.Evaluate(s)
			assert.Nil(t, err)
		})
	}
}
