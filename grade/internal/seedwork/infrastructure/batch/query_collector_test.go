package batch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/testutils"
)

func TestQueryCollectorMultiInsertQuery(t *testing.T) {
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
			q := NewQueryCollector()
			for _, v := range c.params {
				_, err := q.Exec(fmt.Sprintf(sqlTemplate, c.sql), v...)
				assert.Nil(t, err)
			}
			s := testutils.NewDbSessionStub(testutils.NewRowsStub())
			_, err := q.Evaluate(s)
			assert.Equal(t, fmt.Sprintf(sqlTemplate, c.expectedSql), s.ActualQuery)
			assert.Equal(t, c.expectedParams, s.ActualParams)
			assert.Nil(t, err)
		})
	}
}

func TestQueryCollectorAutoincrementMultiInsertQuery(t *testing.T) {
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
			q := NewQueryCollector()
			for _, v := range c.params {
				_, err := q.Exec(fmt.Sprintf(sqlTemplate, c.sql), v...)
				assert.Nil(t, err)
			}
			s := testutils.NewDbSessionStub(testutils.NewRowsStub([]any{1}, []any{2}, []any{3}))
			_, err := q.Evaluate(s)
			assert.Equal(t, fmt.Sprintf(sqlTemplate, c.expectedSql), s.ActualQuery)
			assert.Equal(t, c.expectedParams, s.ActualParams)
			assert.Nil(t, err)
		})
	}
}
